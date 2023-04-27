package service

import (
	"fabric-go-sdk-sample/sdkInit"
	"fmt"
	"os"
	"strings"
)

const (
	cc_name    = "samplecc"
	cc_version = "1.0.0"
)

var App sdkInit.Application

func NewService(config string) {
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "./fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org2",
			OrgMspId:      "Org2MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "./fixtures/channel-artifacts/Org2MSPanchors.tx",
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    "./fixtures/channel-artifacts/mychannel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    "./chaincode/",
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkInit.Setup(config, &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	if err := info.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk); err != nil {

		fmt.Println("InitService successful")
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	App = sdkInit.Application{
		SdkEnvInfo: &info,
	}
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
