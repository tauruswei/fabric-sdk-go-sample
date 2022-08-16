package main

import (
	"brilliance/fast-deploy/chaincode/gconfig/handler"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("G_Config")
var handlers = map[string]handler.Function{
	//gconfig
	"SetConfig":         handler.SetConfig,
	"GetConfig":         handler.GetConfig,
	"SetNetworkOrgs":    handler.SetOrganization,
	"GetAllNetworkOrgs": handler.GetAllOrganizations,
	"GetNetworkOrgs":    handler.GetOrganizations,
	//gevent
	"Notify":    handler.Notify,
	"GetNotify": handler.GetNotify,
	//fileserver
	"SetFileInfo": handler.SetFileInfo,
	"GetFileInfo": handler.GetFileInfo,

	// cert
	"UploadCert":            handler.UploadCert,
	"UploadCerts":           handler.UploadCerts,
	"QueryCertByCommonName": handler.QueryCertByCommonName,
	"QueryAllCert":          handler.QueryAllCert,
}

type GConfigChaincode struct {
}

func (c *GConfigChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	logger.Infof("Init [%s]，args [%s]\n", function, args)
	return shim.Success(nil)
}

func (c *GConfigChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()
	logger.Infof("Invoke [%s] %s \n", function)
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("Invoking function %s err: %s!", function, err)
		}
	}()

	h, ok := handlers[function]
	if !ok {
		logger.Errorf("handlers [%s] faild!!!", function)
		return shim.Error("Function handler faild!!!")
	}

	res, err := h(stub, args)
	if err != nil {
		logger.Errorf("Invoke error, err: %s", err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(res)
}

func main() {
	err := shim.Start(new(GConfigChaincode))
	if err != nil {
		logger.Errorf("Error starting chaincode 'G_Config', err: %s", err)
	}
}
