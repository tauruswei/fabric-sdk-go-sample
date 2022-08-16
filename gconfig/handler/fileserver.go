package handler

import (
	"brilliance/fast-deploy/chaincode/gconfig/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func SetFileInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("SetConfig  args is %#v \n", args)
	if len(args) != 1 {
		return nil, fmt.Errorf("One parameters are required ")
	}
	fileBasic := &model.FileBasic{}
	err := json.Unmarshal([]byte(args[0]), fileBasic)
	if err != nil {
		return nil, fmt.Errorf("%s ：[%s]", "反序列化文件参数报文出错", err.Error())
	}
	if len(fileBasic.Name) == 0 {
		return nil, fmt.Errorf("文件名不能为空！")
	}
	if len(fileBasic.Path) == 0 {
		return nil, fmt.Errorf("文件路径不能为空！")
	}
	if len(fileBasic.Uploader) == 0 {
		return nil, fmt.Errorf("文件上传者不能为空！")
	}
	fmt.Printf("SetFileInfo  fileBasic is %#v \n", fileBasic)

	txid := stub.GetTxID()
	//查询是否已存在key为txid的记录
	kValue, err := stub.GetState(txid)
	if err != nil {
		return nil, fmt.Errorf("GetState  error is : %s!", err.Error())
	}
	if kValue != nil {
		//key为txid的记录已存在，重新上传文件产生新的txid
		return nil, fmt.Errorf("请重新上传文件！")
	}
	//文件信息上链
	byteData, err := json.Marshal(fileBasic)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(txid, byteData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GetFileInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("GetFileInfo  args [%#v]\n", args)
	if len(args) != 1 {
		return nil, fmt.Errorf("One parameters are required ")
	}
	if len(args[0]) == 0 {
		return nil, fmt.Errorf("查询文件参数不能为空！")
	}
	fileBasic, err := stub.GetState(args[0])
	if err != nil {
		return nil, err
	}
	return fileBasic, nil
}
