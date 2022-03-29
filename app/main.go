package main

import (
	"fabric-sdk-go-sample/cli"
	"log"
)

const (
	cfgPath = "/Users/fengxiaoxiao/work/go-projects/fabric-sdk-go-sample/config/config.yaml"
)

var (
	peer0Org1 = "peer0.org1.example.com"
	peer1Org1 = "peer1.org1.example.com"
)

func main() {
	org1Client := cli.New(cfgPath, "Org1", "Admin", "Admin")
	defer org1Client.Close()
	// Install, instantiate, invoke, query
	Phase(org1Client)
}

func Phase(cli1 *cli.Client) {
	log.Println("=================== Phase 1 begin ===================")
	defer log.Println("=================== Phase 1 end ===================")

	//if err := cli1.InstallCC("v1", peer0Org1); err != nil {
	//	log.Panicf("Intall chaincode error: %v", err)
	//}
	//log.Println("Chaincode has been installed on org1's peers")

	// InstantiateCC chaincode only need once for each channel
	//if _, err := cli1.InstantiateCC("v1", peer0Org1); err != nil {
	//	log.Panicf("Instantiated chaincode error: %v", err)
	//}
	//log.Println("Chaincode has been instantiated")

	if _, err := cli1.InvokeCC([]string{peer0Org1}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	//if err := cli1.QueryCC("peer0.org1.example.com", "a"); err != nil {
	//	log.Panicf("Query chaincode error: %v", err)
	//}
	//log.Println("Query chaincode success on peer0.org1")
}
