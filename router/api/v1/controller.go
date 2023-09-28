package v1

import (
	"crypto/sha256"
	"encoding/base64"
	"fabric-go-sdk-sample/log"
	"fabric-go-sdk-sample/model"
	"fabric-go-sdk-sample/result"
	"fabric-go-sdk-sample/sdkInit"
	"fabric-go-sdk-sample/service"
	"fabric-go-sdk-sample/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	cc_name    = "samplecc"
	cc_version = "1.0.0"
)

var logger = log.MustGetLogger("controller-loggger")

func Init(c *gin.Context) {
	g := result.Gin{C: c}

	// init orgs information

	projectRootPath, err := os.Getwd()

	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: projectRootPath + "/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org2",
			OrgMspId:      "Org2MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: projectRootPath + "/fixtures/channel-artifacts/Org2MSPanchors.tx",
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    projectRootPath + "/fixtures/channel-artifacts/mychannel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    projectRootPath + "/chaincode/",
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		logger.Error(result.SERVER_ERROR.FillArgs(">> Create channel and join error:" + err.Error()))
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		logger.Error(result.SERVER_ERROR.FillArgs(">> create chaincode lifecycle error:" + err.Error()))
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	if err := info.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk); err != nil {
		logger.Error(result.SERVER_ERROR.FillArgs(">> InitService unsuccessful error:" + err.Error()))
		os.Exit(-1)
	}

	service.App = sdkInit.Application{
		SdkEnvInfo: &info,
	}

	g.Success("")
}

func Invoke(c *gin.Context) {
	g := result.Gin{C: c}

	request := model.InvokeRequest{}

	err := util.Validator(c, &request, g)
	if err != nil {
		g.Error(result.PARAMETER_VALID_ERROR.FillArgs(err.Error()))
		return
	}

	sum := sha256.Sum256([]byte(request.Data))
	encoded := base64.StdEncoding.EncodeToString(sum[:])

	a := []string{"set", request.Id, encoded}
	ret, err := service.App.Set(a)
	if err != nil {
		logger.Error(result.SERVER_ERROR.FillArgs(err.Error()))
		return
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	g.Success("")
}

func Query(c *gin.Context) {
	g := result.Gin{C: c}

	request := model.InvokeRequest{}

	err := util.Validator(c, &request, g)
	if err != nil {
		g.Error(result.PARAMETER_VALID_ERROR.FillArgs(err.Error()))
		return
	}

	a := []string{"get", request.Id}
	response, err := service.App.Get(a)
	if err != nil {
		logger.Error(result.SERVER_ERROR.FillArgs(err.Error()))
		return
	}
	fmt.Println("<--- 查询信息　--->：", response)

	g.Success(response)
}
