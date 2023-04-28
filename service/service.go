package service

import (
	"fabric-go-sdk-sample/sdkInit"
	"strings"
)

const (
	cc_name    = "samplecc"
	cc_version = "1.0.0"
)

var App sdkInit.Application

func NewService(config string) {
	//// init orgs information
	//
	//orgs := []*sdkInit.OrgInfo{
	//	{
	//		OrgAdminUser:  "Admin",
	//		OrgName:       "Org1",
	//		OrgMspId:      "Org1MSP",
	//		OrgUser:       "User1",
	//		OrgPeerNum:    1,
	//		OrgAnchorFile: "/Users/fengxiaoxiao/work/go-projects/fabric-sdk-go-sample/fixtures/channel-artifacts/Org1MSPanchors.tx",
	//	},
	//	{
	//		OrgAdminUser:  "Admin",
	//		OrgName:       "Org2",
	//		OrgMspId:      "Org2MSP",
	//		OrgUser:       "User1",
	//		OrgPeerNum:    1,
	//		OrgAnchorFile: "/Users/fengxiaoxiao/work/go-projects/fabric-sdk-go-sample/fixtures/channel-artifacts/Org2MSPanchors.tx",
	//	},
	//}
	//
	//// init sdk env info
	//info := sdkInit.SdkEnvInfo{
	//	ChannelID:        "mychannel",
	//	ChannelConfig:    "/Users/fengxiaoxiao/work/go-projects/fabric-sdk-go-sample/fixtures/channel-artifacts/mychannel.tx",
	//	Orgs:             orgs,
	//	OrdererAdminUser: "Admin",
	//	OrdererOrgName:   "OrdererOrg",
	//	OrdererEndpoint:  "orderer.example.com",
	//	ChaincodeID:      cc_name,
	//	ChaincodePath:    "/Users/fengxiaoxiao/work/go-projects/fabric-sdk-go-sample/chaincode/",
	//	ChaincodeVersion: cc_version,
	//}
	//
	//// sdk setup
	//sdk, err := sdkInit.Setup("config.yaml", &info)
	//if err != nil {
	//	fmt.Println(">> SDK setup error:", err)
	//	os.Exit(-1)
	//}
	//
	//// create channel and join
	//if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
	//	fmt.Println(">> Create channel and join error:", err)
	//	os.Exit(-1)
	//}
	//
	//// create chaincode lifecycle
	//if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
	//	fmt.Println(">> create chaincode lifecycle error: %v", err)
	//	os.Exit(-1)
	//}
	//
	//// invoke chaincode set status
	//fmt.Println(">> 通过链码外部服务设置链码状态......")
	//
	//if err := info.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk); err != nil {
	//
	//	fmt.Println("InitService successful")
	//	os.Exit(-1)
	//}
	//
	//App = sdkInit.Application{
	//	SdkEnvInfo: &info,
	//}
}

func computeUpdate(old, new map[string]string) map[string]string {
	update := make(map[string]string)
	for k, v := range new {
		if ov, ok := old[k]; ok {
			if v == ov {
				continue
			}
		}
		update[k] = v
	}
	return update
}

func combinStr(strs []string, sep string) string {
	return strings.Trim(strings.Join(strs, sep), sep)
}
