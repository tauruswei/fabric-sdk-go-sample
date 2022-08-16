package model

const (
	UploadCertObjectType = "UploadCertObj"
	UploadCertEvent      = "UploadCertEvent"
)

const (
	Update = 1
)

type UploadCert struct {
	MspId    string `json:"msp_id"`
	Cert     string `json:"cert"`
	CertType string `json:"cert_type"` // msp tls
	NodeType string `json:"node_type"` // orderer peer
}

type SerialNumber struct {
	// 用作上链处理对象
	ObjectType    string `json:"docType"`
	Serial_Number string `json:"serial_number"`
	IsRootCert    bool   `json:"is_root_cert"`
	MspId         string `json:"msp_id"`
	CommonName    string `json:"common_name"`
	CertType      string `json:"cert_type"` // msp tls
	NodeType      string `json:"node_type"` // orderer peer
}

type CertInfo struct {
	Cert       string `json:"cert"`
	Status     int    `json:"status"`
	Uploadtime int64  `json:"uploadtime"`
}

type QueryCert struct {
	CommonName string `json:"common_name"`
	MspId      string `json:"msp_id"`
	CertType   string `json:"cert_type"` // msp tls
	NodeType   string `json:"node_type"` // orderer peer
}

type NodeHistoryInfo struct {
	CommonName   string     `json:"common_name"`
	CertInfoList []CertInfo `json:"cert_info_list"`
}

type SerialNumberList struct {
	IsRootCert bool   `json:"is_root_cert"`
	MspId      string `json:"msp_id"`
	CommonName string `json:"common_name"`
	CertType   string `json:"cert_type"` // msp tls
	NodeType   string `json:"node_type"` // orderer peer
	//CertInfoList []CertInfo `json:"cert_info_list"`
	SerialNumbers []*SerialNumber `json:"serial_numbers"`
}

type CertList struct {
	IsRootCert   bool       `json:"is_root_cert"`
	MspId        string     `json:"msp_id"`
	CommonName   string     `json:"common_name"`
	CertType     string     `json:"cert_type"` // msp tls
	NodeType     string     `json:"node_type"` // orderer peer
	CertInfoList []CertInfo `json:"cert_info_list"`
}
