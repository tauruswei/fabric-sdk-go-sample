package main

import (
	"fabric-sdk-go-sample/cli"
	"fmt"
	"log"
)

const (
	//cfgPath = "/opt/go-projects/fabric-sdk-go-sample/config/config.yaml"
	cfgPath = "/opt/go/src/github.com/hyperledger/fabric-sdk-go-sample/config/config.yaml"
)

var (
	peer0Org1 = "peer0.org1.example.com"
)

func main() {
	org1Client := cli.New(cfgPath, "Org1", "Admin", "Admin")
	defer org1Client.Close()
	// Install, instantiate, invoke, query
	//Phase1(org1Client)
	//CreateAsset(org1Client)
	//CreateAssets(org1Client)
	QueryAsset(org1Client)
	//QueryAssets(org1Client)
}

func Phase(cli1 *cli.Client) {
	log.Println("=================== Phase 1 begin ===================")
	defer log.Println("=================== Phase 1 end ===================")

	if err := cli1.InstallCC("v1", peer0Org1); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	// InstantiateCC chaincode only need once for each channel
	if _, err := cli1.InstantiateCC("v1", peer0Org1); err != nil {
		log.Panicf("Instantiated chaincode error: %v", err)
	}
	log.Println("Chaincode has been instantiated")

	if _, err := cli1.InvokeCC([]string{peer0Org1}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if err := cli1.QueryCC("peer0.org1.example.com", "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")
}
func QueryAsset(cli1 *cli.Client) {
	log.Println("=================== QueryAsset ===================")
	defer log.Println("=================== QueryAsset ===================")

	if err := cli1.QueryAsset("peer0.org1.example.com", "1234567890123456789012345678901234567890123456789012345678901234"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")
}
func QueryAssets(cli1 *cli.Client) {
	log.Println("=================== QueryOneAsset ===================")
	defer log.Println("=================== QueryOneAsset ===================")

	if err := cli1.QueryAssets("peer0.org1.example.com", "asset1", "asset11", "10"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")
}
func CreateAsset(cli1 *cli.Client) {
	log.Println("=================== CreateAssets ===================")
	defer log.Println("=================== CreateAssets ===================")
	id := "1234567890123456789012345678901234567890123456789012345678901234"
	if _, err := cli1.CreateAsset([]string{peer0Org1}, id); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")
}
func CreateAssets(cli1 *cli.Client) {
	log.Println("=================== CreateAssets ===================")
	defer log.Println("=================== CreateAssets ===================")
	for i := 101; i < 300; i++ {
		id := fmt.Sprintf("asset%d", i)
		if _, err := cli1.CreateAsset([]string{peer0Org1}, id); err != nil {
			log.Panicf("Query chaincode error: %v", err)
		}
		log.Println("Query chaincode success on peer0.org1")
	}
}
