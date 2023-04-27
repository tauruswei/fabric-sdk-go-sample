package global

import (
	"github.com/rogpeppe/go-internal/cache"

	"fabric-go-sdk-sample/config"
	//"github.com/hyperledger/fabric/bccsp/verifier"
	//"github.com/jinzhu/gorm"
	//"github.com/tjfoc/gmsm/sm2"
	"gorm.io/gorm"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/21 下午4:43
 */
// Settings
var (
	// SQLDB 数据库
	SQLDB *gorm.DB
	//Verifier  *verifier.BccspCryptoVerifier
	//Cert      *sm2.Certificate

	OssConfig *config.OssConfig

	SmsConfig *config.SmsConfig
	Cache     *cache.Cache

	CertBytes []byte
)

const (
	// 0 已经发出去，新建的，未查看
	// 1 是已经查看
	CONTRACT_VALID     = 2 // 双方都同意，但是还没有铸造 nft
	CONTRACT_NFT_VALID = 3 // 双方都同意，已经铸造 nft
	CONTRACT_INVALID   = 4 // 双方不同意
	CONTRACT_ALL       = 5 // 所有的合同

	NFT_NOT_ON_LIST = 0
	NFT_ON_SALE     = 1
	NFT_SOLD        = 2

	ORDER_NOT_PAID = 1
	ORDER_PAID     = 2
	ORDER_CANCELED = 3

	CopyrightAuthorizedContract  = "CopyrightAuthorizedContract"
	CopyrightTransactionContract = "CopyrightTransactionContract"
	MusicianCooperationContract  = "MusicianCooperationContract"

	CONTACT_UNHANDLE = 0 // contact 未处理
	CONTACT_HANDLED  = 1 // contact 已处理

)
