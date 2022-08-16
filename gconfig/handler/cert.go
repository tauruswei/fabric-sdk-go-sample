package handler

import (
	"brilliance/fast-deploy/chaincode/gconfig/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"crypto/x509"
	"github.com/tjfoc/gmsm/x509"
	"time"
)

func UploadCert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Printf("UploadCert  args is %#v \n", args)

	// 解析参数
	var data string
	if len(args) > 0 {
		data = args[0]
	}
	uploadCert := model.UploadCert{}
	err := json.Unmarshal([]byte(data), &uploadCert)
	if err != nil {
		return nil, err
	}

	txid, err := uploadOneCert(stub, uploadCert)
	if err != nil {
		return nil, err
	}
	return txid, nil
}

func QueryCertByCommonName(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// 根据证书域名查询所有证书
	// 解析参数
	var data string
	if len(args) > 0 {
		data = args[0]
	}
	queryCert := model.QueryCert{}
	err := json.Unmarshal([]byte(data), &queryCert)
	if err != nil {
		return nil, err
	}
	fmt.Printf(" %c[%d;%d;%dm ==== QueryCertByCommonName data is [%s] ======%c[0m \n", 0x1B, 1, 40, 33, data, 0x1B)
	key, err := stub.CreateCompositeKey(model.UploadCertObjectType, []string{queryCert.NodeType, queryCert.CertType, queryCert.MspId, queryCert.CommonName})
	if err != nil {
		fmt.Printf("CreateCompositeKey failed is %s \n", err)
		return nil, err
	}
	values, err := stub.GetState(key)
	if err != nil {
		fmt.Printf("GetState failed is %s \n", err)
		return nil, err
	}

	var snList []*model.SerialNumber
	if len(values) != 0 {
		err = json.Unmarshal(values, &snList)
		if err != nil {
			fmt.Printf(" %c[%d;%d;%dm ==== Unmarshal failed is %s ======%c[0m \n", 0x1B, 1, 40, 33, err, 0x1B)
			return nil, err
		}
	}
	fmt.Printf(" %c[%d;%d;%dm ==== QueryCertByCommonName len(snList) is [%d] ; snList is [%+v] ======%c[0m \n", 0x1B, 1, 40, 33, len(snList), snList, 0x1B)

	var cInfo []model.CertInfo
	for _, sn := range snList {
		value, err := stub.GetState(sn.Serial_Number)
		if err != nil {
			return nil, err
		}
		cert := model.CertInfo{}
		err = json.Unmarshal(value, &cert)
		if err != nil {
			fmt.Printf(" %c[%d;%d;%dm ==== Unmarshal CertInfo failed is %s ======%c[0m \n", 0x1B, 1, 40, 33, err, 0x1B)
			return nil, err
		}
		cInfo = append(cInfo, cert)
	}

	certHisList := model.NodeHistoryInfo{
		CommonName:   queryCert.CommonName,
		CertInfoList: cInfo,
	}
	fmt.Printf(" %c[%d;%d;%dm ==== QueryCertByCommonName is [%v] ======%c[0m \n", 0x1B, 1, 40, 33, certHisList, 0x1B)
	bytes, err := json.Marshal(certHisList)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Str2Stamp(str string) (int64, error) {
	t, err := time.Parse(defaultLayout, str)
	if err != nil {
		return 0, err
	}
	return t.UnixNano(), nil
}

const (
	defaultLayout = "2006-01-02-15.04.05.000000"
)

func QueryAllCert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// 解析参数
	var data string
	if len(args) > 0 {
		data = args[0]
	}
	fmt.Printf(" %c[%d;%d;%dm ==== QueryAllCert data is [%s] ======%c[0m \n", 0x1B, 1, 40, 33, data, 0x1B)
	queryCert := model.QueryCert{}
	err := json.Unmarshal([]byte(data), &queryCert)
	if err != nil {
		return nil, err
	}
	fmt.Printf(" %c[%d;%d;%dm ==== queryCert.NodeType  is [%s] ;  queryCert.CertType is [%s]======%c[0m \n", 0x1B, 1, 40, 33, queryCert.NodeType, queryCert.CertType, 0x1B)
	iterator, err := stub.GetStateByPartialCompositeKey(model.UploadCertObjectType, []string{queryCert.NodeType, queryCert.CertType})
	defer iterator.Close()
	if err != nil {
		return nil, err
	}

	var snList []model.SerialNumberList
	for iterator.HasNext() {
		var sn []*model.SerialNumber
		item, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		if len(item.Value) > 0 {
			fmt.Printf(" %c[%d;%d;%dm ==== iterator.HasNext()item.Value is [%s] ======%c[0m \n", 0x1B, 1, 40, 33, string(item.Value), 0x1B)
			err = json.Unmarshal(item.Value, &sn)
			var snlist model.SerialNumberList
			snlist.NodeType = sn[0].NodeType
			snlist.CertType = sn[0].CertType
			snlist.MspId = sn[0].MspId
			snlist.CommonName = sn[0].CommonName
			snlist.IsRootCert = sn[0].IsRootCert
			snlist.SerialNumbers = sn
			snList = append(snList, snlist)
		}
		fmt.Printf(" %c[%d;%d;%dm ====QueryAllCert sn is [%+v] ======%c[0m \n", 0x1B, 1, 40, 33, sn, 0x1B)
	}
	var certList []model.CertList
	for _, sns := range snList {
		var certs []model.CertInfo
		for _, sn := range sns.SerialNumbers {
			value, err := stub.GetState(sn.Serial_Number)
			if err != nil {
				return nil, err
			}
			cert := model.CertInfo{}
			err = json.Unmarshal(value, &cert)
			if err != nil {
				return nil, err
			}
			certs = append(certs, cert)
		}
		var cert model.CertList
		cert.CertType = sns.CertType
		cert.NodeType = sns.NodeType
		cert.MspId = sns.MspId
		cert.CommonName = sns.CommonName
		cert.IsRootCert = sns.IsRootCert
		cert.CertInfoList = certs
		certList = append(certList, cert)
	}
	fmt.Printf(" %c[%d;%d;%dm ==== QueryAllCert len(certList) is [%d] ; certList is [%+v] ======%c[0m \n", 0x1B, 1, 40, 32, len(certList), certList, 0x1B)
	certListBytes, err := json.Marshal(certList)
	if err != nil {
		return nil, err
	}
	return certListBytes, nil
}

func uploadOneCert(stub shim.ChaincodeStubInterface, uploadCert model.UploadCert) ([]byte, error) {

	//fmt.Printf(" %c[%d;%d;%dm ==== [UploadCert] uploadCert.MspId is [%s] ; creator.MspId is [%s]======%c[0m \n", 0x1B, 1, 40, 33,uploadCert.MspId ,creator.MspId, 0x1B)
	// 判断如果不是ca证书（自签名）证书，那么则验证签发关系
	// 判断是否是当前组织，只有当前组织才有权限上传本组织证书
	//if creator.MspId != uploadCert.MspId {
	//	return nil, errors.New("非本组织不可上传，上传身份为:" + uploadCert.MspId + "; creator.MspId :"+creator.MspId )
	//}
	//certDERBlock, _ := pem.Decode([]byte(uploadCert.Cert))
	//if certDERBlock == nil {
	//	fmt.Errorf("UploadCert  pem Decode failed , certDERBlock is nil \n")
	//	return nil, fmt.Errorf("UploadCert  pem Decode failed , certDERBlock is nil \n")
	//}
	//cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	cert, err := x509.ReadCertificateFromPem([]byte(uploadCert.Cert))
	if err != nil {
		fmt.Printf("UploadCert  ParseCertificate failed is %s \n", err)
		return nil, err
	}
	var isRootCert bool
	isRootCert = true
	if cert.Issuer.CommonName != cert.Subject.CommonName {
		key, err := stub.CreateCompositeKey(model.UploadCertObjectType, []string{uploadCert.NodeType, uploadCert.CertType, uploadCert.MspId, cert.Issuer.CommonName})
		if err != nil {
			fmt.Printf("CreateCompositeKey failed is %s \n", err)
			return nil, err
		}
		fmt.Printf(" %c[%d;%d;%dm ==== 验证证书签发关系 CreateCompositeKey is [%s] ======%c[0m \n", 0x1B, 1, 40, 33, key, 0x1B)
		// todo 验证证书签发关系
		isRootCert = false
	}

	key, err := stub.CreateCompositeKey(model.UploadCertObjectType, []string{uploadCert.NodeType, uploadCert.CertType, uploadCert.MspId, cert.Subject.CommonName})
	if err != nil {
		fmt.Printf("CreateCompositeKey failed is %s \n", err)
		return nil, err
	}
	values, err := stub.GetState(key)
	if err != nil {
		fmt.Printf("GetState failed is %s \n", err)
		return nil, err
	}
	var snList []*model.SerialNumber
	if len(values) != 0 {
		err = json.Unmarshal(values, &snList)
		if err != nil {
			fmt.Printf("Unmarshal failed is %s \n", err)
			return nil, err
		}
	}

	fmt.Printf(" %c[%d;%d;%dm ==== UploadCert before snList [%v] ======%c[0m \n", 0x1B, 1, 40, 32, snList, 0x1B)
	snMap := make(map[string]*model.SerialNumber)
	for _, s := range snList {
		if _, ok := snMap[s.Serial_Number]; !ok {
			snMap[s.Serial_Number] = s
		}
	}
	if len(snList) > 0 {
		if _, ok := snMap[cert.SerialNumber.String()]; !ok {
			sn := &model.SerialNumber{
				ObjectType:    model.UploadCertObjectType,
				MspId:         uploadCert.MspId,
				NodeType:      uploadCert.NodeType,
				CertType:      uploadCert.CertType,
				CommonName:    cert.Subject.CommonName,
				Serial_Number: cert.SerialNumber.String(),
				IsRootCert:    isRootCert,
			}
			snList = append(snList, sn)
		}
	} else {
		sn := &model.SerialNumber{
			ObjectType:    model.UploadCertObjectType,
			MspId:         uploadCert.MspId,
			NodeType:      uploadCert.NodeType,
			CertType:      uploadCert.CertType,
			CommonName:    cert.Subject.CommonName,
			Serial_Number: cert.SerialNumber.String(),
			IsRootCert:    isRootCert,
		}
		snList = append(snList, sn)
	}

	fmt.Printf(" %c[%d;%d;%dm ==== UploadCert after snList [%v] ======%c[0m \n", 0x1B, 1, 40, 33, snList, 0x1B)
	snbytes, err := json.Marshal(snList)
	if err != nil {
		fmt.Printf("Marshal failed is %s \n", err)
		return nil, err
	}
	err = stub.PutState(key, snbytes)
	if err != nil {
		fmt.Printf("PutState failed is %s \n", err)
		return nil, err
	}
	timeStampInt, err := Str2Stamp(time.Now().Format(defaultLayout))
	if err != nil {
		fmt.Printf("Str2Stamp failed is %s \n", err)
		return nil, err
	}

	if _, ok := snMap[cert.SerialNumber.String()]; !ok {
		certInfo := model.CertInfo{
			Cert:       uploadCert.Cert,
			Status:     model.Update,
			Uploadtime: timeStampInt,
		}
		certBytes, err := json.Marshal(certInfo)
		if err != nil {
			fmt.Printf("Marshal failed is %s \n", err)
			return nil, err
		}
		err = stub.PutState(cert.SerialNumber.String(), certBytes)
		if err != nil {
			fmt.Printf("PutState failed is %s \n", err)
			return nil, err
		}
	}
	return []byte(stub.GetTxID()), nil
}

func UploadCerts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Printf("UploadCerts  args is %#v \n", args)

	// 解析参数
	var data string
	if len(args) > 0 {
		data = args[0]
	}
	var uploadCerts []model.UploadCert
	err := json.Unmarshal([]byte(data), &uploadCerts)
	if err != nil {
		return nil, err
	}

	for _, cert := range uploadCerts {
		_, err := uploadOneCert(stub, cert)
		if err != nil {
			return nil, err
		}
	}

	return []byte(stub.GetTxID()), nil
}
