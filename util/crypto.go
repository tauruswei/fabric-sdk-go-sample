package util

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fabric-go-sdk-sample/config"
	logger "fabric-go-sdk-sample/log"
	"fabric-go-sdk-sample/model"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/signer"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"

	//"github.com/tjfoc/gmsm/sm2"
	"hash"
	"reflect"
)

var sha256Hash hash.Hash = sha256.New()

func GetSha256Code(s string) string {
	sha256Hash.Reset()
	sha256Hash.Write([]byte(s))
	return fmt.Sprintf("%x", sha256Hash.Sum(nil))
}

// ========== AES CBC ===========
// AES CBC模式加密
func AESEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组密钥
	// NewCipher该函数限制了输入key的长度必须为16、24、32,分别对应AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key) // 分组密钥
	if err != nil {
		logger.Errorf("aes encrypt error: %s", err.Error())
		return nil
	}
	blockSize := block.BlockSize()                              // 获取密钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式 iv: key[:blockSize]
	encrypted = make([]byte, len(origData))                     //创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}

// AES CBC模式解密
func AESDencryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, err := aes.NewCipher(key) // 分组密钥
	if err != nil {
		logger.Errorf("aes dencrypt error: %s", err.Error())
		return nil
	}
	blockSize := block.BlockSize()                              // 获取密钥块长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5Unpadding(decrypted)                       // 去除补全码
	return decrypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 根据密钥和连接进行AES加密
func AESEncryptConnection(key string, conn string) (keyBase64 string, connBase64 string) {
	if len(key) == 0 || len(conn) == 0 {
		logger.Error("连接和密钥不能为空")
		return
	}
	// 密钥base64加密
	keyBase64 = base64.StdEncoding.EncodeToString([]byte(key))
	hash := GetSha256Code(key)                                // 密钥进行hash
	realKey := hash[:32]                                      // 取hash之后的前32位作为获得真正的密钥
	encrypted := AESEncryptCBC([]byte(conn), []byte(realKey)) // 加密
	connBase64 = base64.StdEncoding.EncodeToString(encrypted) // 对加密后的数据再进行base64加密
	return
}

// 根据传进来的私钥和连接(加密后的)解密(AES)得到真正的连接
func AESDecryptConnection(keyBase64 string, connBase64 string) string {
	if len(keyBase64) == 0 || len(connBase64) == 0 {
		logger.Error("连接和密钥不能为空")
		return ""
	}
	originKey, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		logger.Errorf("解析密钥失败：%s", err.Error())
		panic(err)
	}
	originConn, err := base64.StdEncoding.DecodeString(connBase64)
	if err != nil {
		logger.Errorf("解析数据库连接失败,Err：%s", err.Error())
		panic(err)
	}
	hash := GetSha256Code(string(originKey))                     // 获得密钥的hash
	realKey := hash[:32]                                         // 取hash的前32位为真正的密钥
	decryptedConn := AESDencryptCBC(originConn, []byte(realKey)) //解密数据库连接
	if decryptedConn == nil {
		logger.Error("解析数据库连接失败")
		panic("解析数据库连接失败")
	}
	return string(decryptedConn)
}

func CopyFields(des interface{}, source interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(des)
	av := reflect.ValueOf(des)
	bt := reflect.TypeOf(source)
	bv := reflect.ValueOf(source)

	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}

	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		bValue := bv.FieldByName(name)
		if name == "UserType" {
			name = "RoleId"
		}
		f := av.Elem().FieldByName(name)

		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			logger.Infof("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}
func GenMnemonic() (string, error) {
	bip39.SetWordList(wordlists.English)
	entropy, err := bip39.NewEntropy(192)
	if err != nil {
		return "", err
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

func GenKey(mnemonic string) (*ecdsa.PrivateKey, error) {
	seed := bip39.NewSeed(mnemonic, "")
	buf := bytes.NewBuffer(seed)
	pri, err := ecdsa.GenerateKey(Curve, buf) // secp256r1
	// pri, err := ecdsa.GenerateKey(crypto.S256(), buf) // secp256k1 golang 不支持
	if err != nil {
		return nil, err
	}
	return pri, nil
}

func GenKeyPem(mnemonic string) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, "")
	buf := bytes.NewBuffer(seed)
	pri, err := ecdsa.GenerateKey(Curve, buf) // secp256r1
	// pri, err := ecdsa.GenerateKey(crypto.S256(), buf) // secp256k1 golang 不支持
	if err != nil {
		return nil, err
	}
	priPEM, err := utils.PrivateKeyToPEM(pri, nil)
	if err != nil {
		return nil, err
	}

	return priPEM, nil
}

// LoadPrivateKey loads a private key from file in keystorePath
func LoadPrivateKey(keystorePath string) (string, error) {
	var err error

	var rawKey string

	walkFunc := func(path string, info os.FileInfo, err error) error {
		rawKeyByte, err := ioutil.ReadFile(path)
		if strings.HasSuffix(path, "_sk") {
			if err != nil {
				logger.Error(GetErrorStackf(err, "读取密钥失败: path = %s", path))
				return errors.WithMessagef(err, "读取密钥失败: path = %s", path)
			}
			rawKey = string(rawKeyByte)
		}
		return nil
	}

	err = filepath.Walk(keystorePath, walkFunc)
	if err != nil {
		logger.Error(GetErrorStackf(err, "读取密钥失败: keystorePath = %s", keystorePath))
		return "", errors.WithMessagef(err, "读取密钥失败: keystorePath = %s", keystorePath)
	}

	return rawKey, err
}

func ImportPrivateKey(keySize int, key, provider, cryptoType, keyStore string) (priv bccsp.Key, s crypto.Signer,
	err error) {
	//// 生成临时目录
	//keyStore = MakeTempdir()
	//defer os.RemoveAll(keyStore)
	csp, err := config.GetBCCSP(provider, "SHA2", 256)
	if err != nil {
		logger.Error(GetErrorStack(err, "获取 bccsp 实例失败"))
		return nil, nil, errors.WithMessage(err, "获取 bccsp 实例失败")
	}
	block, _ := pem.Decode([]byte(key))

	if block == nil {
		logger.Error(GetErrorStackf(err, "解析私钥失败： privateKey= %s", key))
		return nil, nil, errors.WithMessagef(err, "解析私钥失败： privateKey= %s", key)
	}

	switch strings.ToUpper(cryptoType) {
	case "ECC":
		priv, err = csp.KeyImport(block.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: false})
	default:
		logger.Error(GetErrorStackf(err, "不支持的算法：%s", cryptoType))
		return nil, nil, errors.WithMessagef(err, "不支持的算法：%s", cryptoType)
	}

	if err == nil {
		s, err = signer.New(csp, priv)
		if err != nil {
			logger.Error(GetErrorStack(err, "构建 crypto signer 失败"))
			return nil, nil, errors.WithMessage(err, "构建 crypto signer 失败")

		}
	}
	return
}

// default template for X509 certificates
func X509Template(request *model.CertificateRequest) x509.Certificate {

	// generate a serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	var expiry time.Duration
	if request == nil {
		// set expiry to around 10 years
		expiry = 3650 * 24 * time.Hour
	} else {
		expiry = time.Duration(request.Period) * time.Hour
	}
	// round minute and backdate 5 minutes
	notBefore := time.Now().Round(time.Minute).Add(-5 * time.Minute)

	//basic template to use
	x509 := x509.Certificate{
		SerialNumber:          serialNumber,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(expiry),
		BasicConstraintsValid: true,
	}
	return x509
}

// Additional for X509 subject
func SubjectTemplateAdditional(commonName, org, country, province, locality, orgUnit, streetAddress, postalCode string) pkix.Name {
	name := SubjectTemplate()
	name.CommonName = commonName
	if len(org) >= 1 {
		name.Organization = []string{org}
	}
	if len(country) >= 1 {
		name.Country = []string{country}
	}
	if len(province) >= 1 {
		name.Province = []string{province}
	}

	if len(locality) >= 1 {
		name.Locality = []string{locality}
	}
	if len(orgUnit) >= 1 {
		name.OrganizationalUnit = []string{orgUnit}
	}
	if len(streetAddress) >= 1 {
		name.StreetAddress = []string{streetAddress}
	}
	if len(postalCode) >= 1 {
		name.PostalCode = []string{postalCode}
	}
	return name
}

// default template for X509 subject
func SubjectTemplate() pkix.Name {
	return pkix.Name{
		Country:  []string{"US"},
		Locality: []string{"San Francisco"},
		Province: []string{"California"},
	}
}
