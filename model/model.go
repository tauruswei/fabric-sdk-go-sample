package model

import (
	"encoding/json"
	//"google.golang.org/genproto/googleapis/storage/v1"
	"io"
	"io/ioutil"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:46
 */
/*
	业务请求结构体
*/
type Envelope struct {
	Data        []byte `json:"data"`        // 业务合约请求数据, 对应 QueryBaseInfo
	Sig         []byte `json:"sig"`         // 签名值
	Certificate []byte `json:"certificate"` // 证书
}

/*
 业务合约请求数据  结构体
*/
type QueryBaseInfo struct {
	ContractName string                 `json:"contractName,omitempty"` // 应用合约名字
	Method       string                 `json:"method,omitempty"`       // 请求的方法
	Params       map[string]interface{} `json:"params,omitempty"`       // 请求的数据
	SourceUrl    string                 `json:"sourceUrl,omitempty"`    // 数据源地址
}

func GetBody(body io.ReadCloser, v interface{}) error {

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

type NewCARequest struct {
	CertificateRequest
	KeyRequest
}
type CertificateRequest struct {
	Org           string `json:"org"`
	OrgUnit       string `json:"orgUnit"`
	Country       string `json:"country"`
	CommonName    string `json:"commonName" binding:"required"`
	Province      string `json:"province"`
	Locality      string `json:"locality"`
	StreetAddress string `json:"streetAddress"`
	PostalCode    string `json:"postalCode" binding:"omitempty,email"`
	IssuerSubject string `json:"issuerSubject"`
	IsCA          bool   `json:"isCA"`
	Period        int    `json:"period" binding:"required,gt=0,lte=876000"`
}
type CertificateSigningRequest struct {
	Org           string `json:"org"`
	OrgUnit       string `json:"orgUnit"`
	Country       string `json:"country"`
	CommonName    string `json:"commonName"`
	Province      string `json:"province"`
	Locality      string `json:"locality"`
	StreetAddress string `json:"streetAddress"`
	PostalCode    string `json:"postalCode"`
	CryptoType    string `json:"cryptoType"`
	IsCA          bool   `json:"isCA"`
	KeyName       string `json:"keyName"`
	Period        int    `json:"period"`
	Provider      string `json:"provider"`
}

type SignCertRequest struct {
	NickName     string `json:"nickName"`
	TlsPublicKey string `json:"tlsPublicKey" binding:"required"`
	MspPublicKey string `json:"mspPublicKey" binding:"required"`
	UserType     string `json:"userType"`
	//Passwd       string `json:"passwd" binding:"required,Passwd,max=16,min=8" reg_error_info:"密码必须包含大写字母、小写字母、数字和特殊字符，长度8-16位"`
	//PublicKey string `json:"publicKey" binding:"required"`
}

type KeyRequest struct {
	CryptoType string `json:"cryptoType" binding:"required,oneof=ECC SM2"`
	KeySize    int    `json:"keySize" binding:"required,oneof=256 384"`
	Provider   string `json:"provider" binding:"required"`
}
type RevokeRequest struct {
	CertificateSubject string `json:"certificateSubject"`
}
type CrlRequest struct {
	IssuerSubject string `json:"issuerSubject"`
}

type options struct {
	A string
	B string
	C int
}

var defaultOptions = &options{
	A: "a",
	B: "b",
	C: 0,
}

func NewOption(A, B string, C int) *options {
	return &options{
		A: A,
		B: B,
		C: C,
	}
}
func NewOption2(opts ...Option) *options {
	opt := defaultOptions
	for _, o := range opts {
		o(opt)
	}
	return opt
}

type Option func(*options)

func WithA(A string) Option {
	return func(o *options) {
		o.A = A
	}
}
func WithB(B string) Option {
	return func(o *options) {
		o.B = B
	}
}
func WithC(C int) Option {
	return func(o *options) {
		o.C = C
	}
}

type RegisterUserRequest struct {
	NickName       string `json:"nickName"`
	Name           string `json:"name" `
	IdentityCard   string `json:"identityCard" binding:"required"`
	MobileNumber   string `json:"mobileNumber" binding:"required"`
	RandomCodes    string `json:"randomCodes" binding:"required"`
	EnterpriseName string `json:"enterpriseName"`
}

// todo 添加正则校验
type GenMnemonicsRequest struct {
	NickName string `json:"nickName" binding:"required"`
	Passwd   string `json:"passwd" binding:"required,Passwd,max=16,min=8" reg_error_info:"密码必须包含大写字母、小写字母、数字和特殊字符，长度8-16位"`
	UserType string `json:"userType" binding:"required"`
}

// todo 添加正则校验
type GenRandCodesRequest struct {
	NickName     string `json:"nickName"`
	MobileNumber string `json:"mobileNumber" binding:"required"`
}

// todo 添加正则校验
type LoginRequest struct {
	NickName string `json:"nickName" binding:"required"`
	Passwd   string `json:"passwd" binding:"required"`
}

// todo 添加正则校验
type LogoutRequest struct {
	NickName string `json:"nickName"`
}

// todo 添加正则校验
type ModifyPasswdRequest struct {
	NickName  string `json:"nickName"`
	OldPasswd string `json:"oldPasswd" binding:"required"`
	Newpasswd string `json:"newPasswd" binding:"required,Passwd,max=16,min=8" reg_error_info:"密码必须包含大写字母、小写字母、数字和特殊字符，长度8-16位"`
}

// todo 添加正则校验
type ModifyPersonalDataRequest struct {
	NickName     string `json:"nickName"`
	Name         string `json:"name"`
	IdentityCard string `json:"identityCard"`
	Email        string `json:"email"`
}

// todo 添加正则校验
type ModifyMobileNumberRequest struct {
	NickName        string `json:"nickName"`
	OldMobileNumber string `json:"oldMobileNumber" binding:"required"`
	OldRandomCodes  string `json:"oldRandomCodes" binding:"required"`
	NewMobileNumber string `json:"newMobileNumber" binding:"required"`
	NewRandomCodes  string `json:"newRandomCodes" binding:"required"`
}

// todo 添加正则校验
type UserListRequest struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum" binding:"required"`
}

// todo 添加正则校验
type SearchUserRequest struct {
	NickName string `json:"nickName"`
	Name     string `json:"name"`
	UserType int    `json:"userType"`
	UserListRequest
}

type DataListResult struct {
	BookMark string      `json:"bookMark"`
	Count    int64       `json:"count"`
	DataList interface{} `json:"dataList"`
}

type NFTDataListResult struct {
	BookMark string        `json:"bookMark"`
	Count    int64         `json:"count"`
	DataList []interface{} `json:"dataList"`
}
type CreateNft721Request struct {
	NickName string `json:"nickName"`
	Name     string `json:"name" binding:"required"`
	Label    string `json:"label"`
	Uri      string `json:"uri"`
	Desc     string `json:"desc"`
	Type     int    `json:"type" binding:"required"`
}
type InvokeRequest struct {
	Token string `json:"token"`
}
type SaveNFTToDBRequest struct {
	//OwnerName string `json:"ownerName" binding:"required"`
	//UserType  int    `json:"userType" binding:"required"`
	CreateMusicNFTRequest
}
type SaveContactRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type SellNft721Request struct {
	NickName string  `json:"nickName"`
	TokenId  string  `json:"token_id" binding:"required"`
	Price    float64 `json:"price" binding:"required,gte=0"`
}
type ModifyNft721PriceRequest struct {
	NickName string  `json:"nickName"`
	TokenId  string  `json:"token_id" binding:"required"`
	Price    float64 `json:"price" binding:"required,gte=0"`
}
type NFTListRequest struct {
	NickName string `json:"nickName"`
	//Status   int    `json:"status" binding:"required"` // nft 状态   0 未上架； 1 正在卖； 2 已卖出
	Status   int    `json:"status"` // nft 状态   0 未上架； 1 正在卖； 2 已卖出
	OnSale   bool   `json:"on_sale"`
	Sold     bool   `json:"sold"`
	IsUser   bool   `json:"isUser"` // true 查询用户的 NFT 列表  false  查询平台的 NFT 列表
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
	BookMark string `json:"bookMark"` // chaincode 分页查询参数, 比如 bookmark = "test" ，chaincode查询的时候，会从 bookmark 开始查询
	KeyWord  string `json:"keyword"`  // 关键词搜索
	Label    string `json:"label"`    // 标签分类
}
type ContractListRequest struct {
	NickName string `json:"nickName"`
	UserName string `json:"userName"` // 用户真实姓名
	IsUser   bool   `json:"isUser"`   // true 查询用户的 合同 列表  false  查询平台的 合同 列表
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
	Type     string `json:"type"`     //合同类型
	Status   int    `json:"status"`   // 合同状态 1/2 待生效的合同  3 生效的合同  4 未生效的合同
	BookMark string `json:"bookMark"` // chaincode 分页查询参数, 比如 bookmark = "test" ，chaincode查询的时候，会从 bookmark 开始查询
}
type DelegatedNFTListRequest struct {
	NickName   string `json:"nickName"`
	ContractId string `json:"contractId" binding:"required"` // 合同id
	PageSize   int    `json:"pageSize"`
	PageNum    int    `json:"pageNum"`
	BookMark   string `json:"bookMark"` // chaincode 分页查询参数, 比如 bookmark = "test" ，chaincode查询的时候，会从 bookmark 开始查询
}
type CreateMusicContractRequest struct {
	NickName string `json:"nickName"`
	// todo 判断时间
	//UpdateTime time.Time `form:"update_time" binding:"required,timing" time_format:"2006-01-02 15:04:05"`
	Contract
}
type CreateMusicDelegateContractRequest struct {
	NickName string `json:"nickName"`

	DelegateContract
}
type UpdateContractStatusRequest struct {
	NickName   string   `json:"nickName"`
	ContractId string   `json:"contractId"` // 合同 token id
	MusicIds   []int    `json:"musicIds"`   // 音乐 ids
	DTokenIds  []string `json:"dTokenIds"`  //
	StartTime  []string `json:"startTime"`  // 授权开始时间
	EndTime    []string `json:"endTime"`    // 授权结束时间
	OwnerName  string   `json:"ownerName"`  // 客户/版权方的 昵称
	UserType   int      `json:"userType"`   // 用户类型
}
type CreateMusicNFTRequest struct {
	NickName string `json:"nickName"`
	Music
}
type CreateDelegatedMusicNFTRequest struct {
	NickName     string   `json:"nickName"`
	TokenIds     []string `json:"tokenIds" binding:"required"`           // 原始音乐的tokenIds 或者 授权 NFT 的 tokenIds
	OwnerName    string   `json:"ownerName" binding:"required"`          // 被授权方的 名称
	ContractId   string   `json:"contractId" binding:"required"`         // 合同 token id
	Expiration   []string `json:"expiration"`                            // 过期 时间
	ContractType int      `json:"contractType,omitempty" `               // 合同类型 1 音乐合作  2  版权交易  3版权授权
	UserType     int      `json:"userType,omitempty" binding:"required"` // 用户类型
	MessageId    int      `json:"messageId"`                             // 消息的id
	OwnerAddr    string   `json:"ownerAddr"`                             // 被授权方的地址
}

type QueryContactDetailRequest struct {
	NickName     string `json:"nickName"`
	TokenId      string `json:"tokenId" binding:"required"` // 合同 token id
	ContractType int    `json:"contractType,omitempty" `    // 合同类型 1 音乐合作  2  版权交易  3版权授权
}
type QueryMusicNFTTokenRequest struct {
	NickName   string `json:"nickName"`
	Name       string `json:"name,omitempty" `                        // 歌曲名称
	MusicType  int    `json:"musicType,omitempty" binding:"required"` // 类型：1 歌曲，2 MV；3 微电影
	Singer     string `json:"singer,omitempty"`                       // 表演者名称
	SongWriter string `json:"songWriter,omitempty"`                   // 词作者
	Composer   string `json:"composer,omitempty"`                     // 曲作者
	Player     string `json:"player,omitempty"`                       // 表演者
	Director   string `json:"director,omitempty"`                     // 导演
	Producer   string `json:"producer,omitempty"`                     // 制片人
	PageSize   int    `json:"pageSize"`
	PageNum    int    `json:"pageNum"`
}
type QueryNFTListForMusicianRequest struct {
	NickName   string `json:"nickName"`
	OwnerName  string `json:"ownerName"`            //音乐人或者版权方的名称
	UserType   int    `json:"userType"`             //用户类型 1 管理员  2 版权方  3 普通用户
	Name       string `json:"name,omitempty" `      // 歌曲名称
	MusicType  int    `json:"musicType,omitempty"`  // 类型：1 歌曲，2 MV；3 微电影
	Singer     string `json:"singer,omitempty"`     // 表演者名称
	SongWriter string `json:"songWriter,omitempty"` // 词作者
	Composer   string `json:"composer,omitempty"`   // 曲作者
	Player     string `json:"player,omitempty"`     // 表演者
	Director   string `json:"director,omitempty"`   // 导演
	Producer   string `json:"producer,omitempty"`   // 制片人
	PageSize   int    `json:"pageSize"`
	PageNum    int    `json:"pageNum"`
}

type QueryMusicDelegatedTokenRequest struct {
	NickName        string `json:"nickName"`
	OriginalTokenId string `json:"tokenId" binding:"required"` // 歌曲的 token id
	OwnerName       string `json:"ownerName"`                  // 版权方名称
}

type QueryNFTDetailRequest struct {
	TokenId string `json:"tokenId,omitempty"`
}

type NFT struct {
	TokenId    string `json:"tokenId,omitempty"`
	Name       string `json:"name,omitempty" binding:"required"`
	Label      string `json:"label,omitempty"`
	UriPicture string `json:"uriPicture,omitempty"` // 图片的网址
	UriVideo   string `json:"uriVideo,omitempty"`   // 音频的网址
	Desc       string `json:"desc,omitempty"`
	//Status     int     `json:"status,omitempty" binding:"required"` // nft 状态   0 未上架； 1 正在卖； 2 已卖出
	Status    int     `json:"status,omitempty"` // nft 状态   0 未上架； 1 正在卖； 2 已卖出
	Price     float64 `json:"price,omitempty"`
	Owner     string  `json:"owner,omitempty"`
	OwnerName string  `json:"ownerName,omitempty"` // 所有者名称
}

// 歌曲名称，演唱真，词作者，曲作者，确定歌曲的唯一性
type Music struct {
	NFT
	MusicType               int    `json:"musicType,omitempty"`               //类型：1 歌曲，2 MV；3 微电影
	Number                  string `json:"number,omitempty"`                  // ISRC 号码
	CollectionName          string `json:"collectionName,omitempty"`          // 转接名称
	Singer                  string `json:"singer,omitempty"`                  // 表演者名称
	IssuingDate             string `json:"issuingDate,omitempty"`             // 发行时间
	SongWriter              string `json:"songWriter,omitempty"`              // 词作者
	Composer                string `json:"composer,omitempty"`                // 曲作者
	SongProportion          string `json:"songProportion,omitempty"`          // 词权比例
	ComposerProportion      string `json:"composerProportion,omitempty"`      // 曲权比例
	NeighboringRightsPropor string `json:"neighboringRightsPropor,omitempty"` // 邻接权权利比例
	Language                string `json:"language,omitempty"`                // 语种
	MV
}
type MV struct {
	Player            string `json:"player,omitempty"`            // 表演者
	Director          string `json:"director,omitempty"`          // 导演
	Producer          string `json:"producer,omitempty"`          // 制片人
	Time              string `json:"time,omitempty"`              // 授权方拥有的权利期限
	CopyrightProption string `json:"copyrightProption,omitempty"` // 著作权比例
}
type Contract struct {
	ConmmonContract
	DelegateStartDate     string `json:"delegateStartDate,omitempty"`     // 授权截止日期  2006-01-02 15:04:05
	DelegateEndtDate      string `json:"delegateEndDate,omitempty"`       // 授权开始日期	 2006-01-02 15:04:05
	DelegateRelation      string `json:"delegateRelation,omitempty"`      // 授权关系 代理
	DelegateType          string `json:"delegateType,omitempty"`          // 授权形式 独家/非独家
	DelegateRegion        string `json:"delegateRegion,omitempty"`        // 授权区域 大陆
	DelegateForm          string `json:"delegateForm,omitempty"`          // 授权形式 付费
	CanTransferDelegation bool   `json:"canTransferDelegation,omitempty"` // 是否可以转授权
	Remark                string `json:"remark"`                          // 备注
}
type DelegateContract struct {
	ConmmonContract
	PublicationRight         string `json:"publicationRight,omitempty"`         // 发表权
	SignatureRight           string `json:"signatureRight,omitempty"`           // 署名权
	AmendmentRight           string `json:"amendmentRight,omitempty"`           // 修改权
	KeepIntegrityRight       string `json:"keepIntegrityRight,omitempty"`       // 保护作品完整权
	ReproductionRight        string `json:"reproductionRight,omitempty"`        // 复制权
	IssuingRight             string `json:"issuingRight,omitempty"`             // 发行权
	RentalRight              string `json:"rentalRight,omitempty"`              // 出租权
	ExhibitionRight          string `json:"exhibitionRight,omitempty"`          // 展览权
	PerformingRight          string `json:"performingRight,omitempty"`          // 表演权
	ShowRight                string `json:"showRight,omitempty"`                // 放映权
	BroadcastRight           string `json:"broadcastRight,omitempty"`           // 广播权
	NetworkTransmissionRight string `json:"networkTransmissionRight,omitempty"` // 信息网络传播权
	FilmingRight             string `json:"filmingRight,omitempty"`             // 摄制权
	AdaptRight               string `json:"adaptRight,omitempty"`               // 改编权
	TranslationRight         string `json:"translationRight,omitempty"`         // 翻译权
	CompilationRight         string `json:"compilationRight,omitempty"`         // 汇编权
	OtherRight               string `json:"otherRight,omitempty"`               // 应当由著作权人享有的其他权利
	Remark                   string `json:"remark,omitempty"`
}
type ConmmonContract struct {
	OwnerName      string `json:"ownerName,omitempty" binding:"required"`      // 客户名称
	OtherOwnerName string `json:"otherOwnerName,omitempty" binding:"required"` // 版权方名称
	ContractType   int    `json:"contractType,omitempty" binding:"required"`   // 合同类型 1 音乐合作  2  版权交易  3版权授权
	NFT
}

// todo name owner 需要建立索引
type NFT994 struct {
	TokenId       uint64 `json:"tokenId,omitempty"`       // 授权token的id
	RootTokenId   uint64 `json:"rootTokenId,omitempty"`   // 歌曲的 721token id
	ParentTokenId uint64 `json:"parentTokenId,omitempty"` // parent delegated token
	ContractId    string `json:"contractId,omitempty"`    // 合同的 token id
	Name          string `json:"name,omitempty"`          // 歌曲名称
	OwnerName     string `json:"ownerName,omitempty"`     // 版权方名称
	Owner         string `json:"owner,omitempty"`         // 版权方地址
	Expiration    string `json:"expiration,omitempty"`    // 授权截止日期  format：2006-01-02 15:04:05，该属性应对特例：一个合同可能存在授权不同期限的歌曲
}
type ConnectionConfig struct {
	Name          string                  `json:"name"` // 网络名字
	Version       string                  `json:"version"`
	Client        Client                  `json:"client"`
	Channels      map[string]Channel      `json:"channels"`
	Organizations map[string]Organization `json:"organizations"`
	Orderers      map[string]Node         `json:"orderers"`
	Peers         map[string]Node         `json:"peers"`
}
type Client struct {
	ContractName string     `json:"contractName"` //合约名称
	ChannelName  string     `json:"channelName"`  //通道名称
	Organization string     `json:"organization"` // 当前 sdk 属于哪个组织
	Connection   Connection `json:"connection"`
}
type Connection struct {
	Timeout struct {
		Peer struct {
			Endorder string `json:"endorder"`
		} `json:"peer"`
		Orderer string `json:"orderer"`
	} `json:"timeout"`
}
type Channel struct {
	Orderers []string        `json:"orderers"`
	Peers    map[string]Peer `json:"peers"`
}
type Orderers struct {
}
type Peer struct {
	EndorsingPeer  bool `json:"endorsingPeer"`
	ChaincodeQuery bool `json:"chaincodeQuery"`
	LedgerQuery    bool `json:"ledgerQuery"`
	EventSource    bool `json:"eventSource"`
}
type Organization struct {
	MspId string   `json:"mspid"`
	Peers []string `json:"peers"`
}
type Node struct {
	Url         string `json:"url"`
	Mspid       string `json:"mspid"`
	GrpcOptions struct {
		SslTargetNameOverride string `json:"ssl-target-name-override"`
		HostnameOverride      string `json:"hostnameOverride"`
	} `json:"grpcOptions"`
	TlsCACerts struct {
		Pem string `json:"pem"`
	} `json:"tlsCACerts"`
}

// todo 返回token与usertype
type Loginresponse struct {
	Token    string `json:"token"`
	UserType string `json:"userType"`
	Name     string `json:"name"` // 企业名称或用户真实姓名
}

// todo 返回tlsMnemonics,mspMnemonics
type GenMnemonicsResponce struct {
	TlsMnemonics string `json:"tlsMnemonics"`
	MspMnemonics string `json:"mspMnemonics"`
}

type PaymentsRequest struct {
	TradeType string `json:"tradeType"`
	// AppAuthToken string   `json:"-"`                       // 可选
	// OutTradeNo   string   `json:"out_trade_no,omitempty"`  // 订单支付时传入的商户订单号, 与 TradeNo 二选一
	// TradeNo      string   `json:"trade_no,omitempty"`      // 支付宝交易号
	// QueryOptions []string `json:"query_options,omitempty"` // 可选 查询选项，商户通过上送该字段来定制查询返回信息 TRADE_SETTLE_INFO
}

type QuerypayRequest struct {
	PaymentType    string `json:"paymentType"`
	AppAuthToken   string `json:"-"`                 // 可选
	ProductCode    string `json:"product_code"`      // 必选 业务产品码， 收发现金红包固定为：STD_RED_PACKET； 单笔无密转账到支付宝账户固定为：TRANS_ACCOUNT_NO_PWD； 单笔无密转账到银行卡固定为：TRANS_BANKCARD_NO_PWD
	BizScene       string `json:"biz_scene"`         // 必选 描述特定的业务场景，可传的参数如下： PERSONAL_COLLECTION：C2C现金红包-领红包； DIRECT_TRANSFER：B2C现金红包、单笔无密转账到支付宝/银行卡
	OutBizNo       string `json:"out_biz_no"`        // 可选 商户端的唯一订单号，对于同一笔转账请求，商户需保证该订单号唯一。
	OrderId        string `json:"order_id"`          // 可选 支付宝转账单据号
	PayFundOrderId string `json:"pay_fund_order_id"` // 可选 支付宝支付资金流水号
}

// AddCartRequest 购物车添加商品
type AddCartRequest struct {
	UserID   int    `json:"userId"`
	NftId    int    `json:"nftId"`
	NickName string `json:"nickName"`
}

// CreateOrderRequest 创建订单
type CreateOrderRequest struct {
	UserID int `json:"userId" binding:"required"`
	NftId  int `json:"nftId" binding:"required"`
}

type CreateNFTReleaseRequest struct {
	UrlWin     string `json:"urlWin,omitempty"`
	UrlMac     string `json:"urlMac,omitempty"`
	UrlIos     string `json:"urlIos,omitempty"`
	UrlAndroid string `json:"urlAndroid,omitempty"`
	Version    string `json:"version,omitempty" binding:"required"`
	Comment    string `json:"comment,omitempty"`
}
