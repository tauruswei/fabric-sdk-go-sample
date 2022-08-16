package handler

import (
	"brilliance/fast-deploy/chaincode/gconfig/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("G_Config")

type Function func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error)

func SetConfig(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("SetConfig  args is %#v \n", args)
	if len(args) != 2 {
		return nil, fmt.Errorf("key or value can not empty !!! ")
	}
	creator, err := getCreator(stub)

	data := model.Data{
		MspId: creator.MspId,
		//Value: json.RawMessage(args[1]),
		Value: args[1],
	}
	fmt.Printf("SetConfig  data is %#v \n", data)
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(args[0], byteData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GetConfig(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("GetConfig  args [%#v]\n", args)
	ss, err := stub.GetState(args[0])
	if err != nil {
		return nil, err
	}
	fmt.Printf("stub.GetState(args[0])  [%#v]\n", string(ss))
	return stub.GetState(args[0])
}

func GetHistoryConfig(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("GetHistoryConfig  args [%#v]\n", args)
	iter, err := stub.GetHistoryForKey(args[0])
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	values := make(map[string]model.Data)

	for iter.HasNext() {
		fmt.Printf("next\n")
		if kv, err := iter.Next(); err == nil {
			fmt.Printf("id: %s value: %#v\n", kv.TxId, kv.Value)
			data := model.Data{}
			err = json.Unmarshal(kv.Value, data)
			values[kv.TxId] = data
		}
		if err != nil {
			return nil, err
		}
	}
	bytes, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	return bytes, nil

}

func getCreator(stub shim.ChaincodeStubInterface) (*model.Creator, error) {
	creator_pb, err := stub.GetCreator()
	if err != nil {
		errstr := fmt.Sprintf("get creator from proposal error:%s", err.Error())
		logger.Error(errstr)
		return nil, errors.New(errstr)
	}
	creator := &model.Creator{}
	if err := proto.Unmarshal(creator_pb, creator); err != nil {
		errstr := fmt.Sprintf("unmarshal creator error:%s", err.Error())
		logger.Error(errstr)
		return nil, errors.New(errstr)
	}

	return creator, nil
}
