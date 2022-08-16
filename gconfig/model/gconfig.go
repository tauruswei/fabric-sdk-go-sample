package model

import (
	"github.com/golang/protobuf/proto"
)

const (
	NETWORK_ORGS_KEY = "network_orgs"
)

type Creator struct {
	MspId string `protobuf:"bytes,1,opt,name=mspid" json:"mspid,omitempty"`
}

func (c *Creator) Reset()         { *c = Creator{} }
func (c *Creator) String() string { return proto.CompactTextString(c) }
func (*Creator) ProtoMessage()    {}

type Data struct {
	MspId string `json:"msp_id"`
	//Value json.RawMessage `json:"value"`
	//Value []byte `json:"value"`
	Value string `json:"value"`
}

type Organization struct {
	MspId  string `json:"msp_id"`
	Name   string `json:"name"`
	NameZh string `json:"name_zh"`
}

type NetOrganization struct {
	Peer    []Organization `json:"peer"`
	Orderer []Organization `json:"orderer"`
}
