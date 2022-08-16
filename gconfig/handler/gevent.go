package handler

import (
	"brilliance/fast-deploy/chaincode/gconfig/model"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Notify(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var data string
	if len(args) > 0 {
		data = args[0]
	}
	event := model.NotifyEvent{}
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		return nil, err
	}

	err = stub.PutState(stub.GetTxID(), []byte(data))
	if err != nil {
		return nil, err
	}

	err = stub.SetEvent(event.Name, []byte(data))
	if err != nil {
		return nil, err
	}

	return []byte(stub.GetTxID()), nil
}

func GetNotify(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var txId string
	if len(args) > 0 {
		txId = args[0]
	}
	return stub.GetState(txId)
}
