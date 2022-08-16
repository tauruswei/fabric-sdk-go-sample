package handler

import (
	"brilliance/fast-deploy/chaincode/gconfig/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
)

//orderer组织信息或peer组织信息上链
func SetOrganization(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("SetOrganization  args is %#v \n", args)
	if len(args) != 3 {
		return nil, fmt.Errorf("Incorrect number of parameters !!! ")
	}

	var orgType string
	if strings.Index(args[0], "Orderer") >= 0 {
		orgType = "orderer"
	} else if strings.Index(args[0], "Peer") >= 0 {
		orgType = "peer"
	}

	data := model.Organization{
		MspId:  args[0],
		Name:   args[1],
		NameZh: args[2],
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	orgskey, err := stub.CreateCompositeKey(model.NETWORK_ORGS_KEY, []string{orgType, args[0]})
	fmt.Printf("SetOrganization  orgskey is %#v \n", orgskey)
	err = stub.PutState(orgskey, byteData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

/*
args为空 获取所有的orderer组织与peer组织
*/
func GetAllOrganizations(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("Incorrect number of parameters !!! ")
	}
	iterator, err := stub.GetStateByPartialCompositeKey(model.NETWORK_ORGS_KEY, []string{"peer"})
	defer iterator.Close()
	if err != nil {
		return nil, err
	}
	var peer []model.Organization
	for iterator.HasNext() {
		org := &model.Organization{}
		item, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		if len(item.Value) > 0 {
			err = json.Unmarshal(item.Value, org)
			peer = append(peer, *org)
		}
	}

	ordererIterator, err := stub.GetStateByPartialCompositeKey(model.NETWORK_ORGS_KEY, []string{"orderer"})
	defer ordererIterator.Close()
	if err != nil {
		return nil, err
	}
	var orderer []model.Organization
	for ordererIterator.HasNext() {
		org := &model.Organization{}
		item, err := ordererIterator.Next()
		if err != nil {
			return nil, err
		}
		if len(item.Value) > 0 {
			err = json.Unmarshal(item.Value, org)
			orderer = append(orderer, *org)
		}
	}

	data := model.NetOrganization{
		Peer:    peer,
		Orderer: orderer,
	}
	fmt.Printf("GetOrganizations  return data is %#v \n", data)
	orgsByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return orgsByte, nil
}

/*
args[0] = "orderer" 获取所有的orderer组织
args[0] = "orderer"  args[1] = "citicOrderer"  获取citicOrderer的orderer组织信息
args[0] = "peer" 获取所有的peer组织
args[0] = "peer"  args[1] = "citicPeer"  获取citicPeer的peer组织信息
*/
func GetOrganizations(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("Required input parameters !!! ")
	}

	iterator, err := stub.GetStateByPartialCompositeKey(model.NETWORK_ORGS_KEY, args)
	defer iterator.Close()
	if err != nil {
		return nil, err
	}
	var data []model.Organization
	for iterator.HasNext() {
		org := &model.Organization{}
		item, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		if len(item.Value) > 0 {
			err = json.Unmarshal(item.Value, org)
			data = append(data, *org)
		}
	}
	fmt.Printf("GetOrganizations  return data is %#v \n", data)
	orgsByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return orgsByte, nil
}
