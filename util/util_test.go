package util

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
	"math/big"
	logger "nft/nft-marketplace/common/log"
	"strconv"
	"testing"
	"time"
)

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2021/12/7 11:14 上午
 */
func TestGenMnemonicGenKey(t *testing.T) {
	mnemonic, _ := GenMnemonic()
	fmt.Println(mnemonic)
	key, err := GenKey(mnemonic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key.PublicKey.X)
	fmt.Println(key.PublicKey.Y)

}
func TestGenKey(t *testing.T) {
	str := "牙 泵 死 孩 步 反 先 差 徒 胺 落 狱"
	//牙 泵 死 孩 步 反 先 差 徒 胺 落 狱 校 自 称 胆 炼 还 改 燃 戴 猛 况
	//即 沉 霍 哀 质 绳 候 范 碰 住 警 症
	key, err := GenKey(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key.PublicKey.X)
	fmt.Println(key.PublicKey.Y)
}

func TestUint64Rand(t *testing.T) {
	var str []string
	rand := Uint64Rand()
	fmt.Println(rand)

	formatUint := strconv.FormatUint(rand, 10)

	for i := 0; i < 1000000; i++ {
		str = append(str, formatUint)
	}
	marshal, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(marshal))

	file, err := ioutil.ReadFile("/Users/fengxiaoxiao/Desktop/src.tar.gz")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(file))

}
func TestParsePrivateKey(t *testing.T) {
	//raw:="-----BEGIN PRIVATE KEY-----\nMIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQgAtYUQXpdwrO+oDvCWZqyTqowOpg9U6Lg7x8PBjyEjh+hRANCAARx5BkVCfdKmMCCyoDnwK2bsP6+2SuN7ryKkIcVG5eQhd8NkmndvAuhvGBpmx5DQbjb+r9coUFKEuGxsSNu0UH6\n-----END PRIVATE KEY-----"
	//raw:="-----BEGIN PRIVATE KEY-----\nMHcCAQEEIGqxQ3Bkd7kge8FJV02vZ/NN5+99JatjIG13cVJ8bfV3oAoGCCqGSM49AwEHoUQDQgAEVlhWOXdUZkrh4ns49SNV1OjlYFomf4jNYyUvR4XFifyNGHJlfCzHqKnbhQMl7GsYyZTnAEr+QFOha1+7dBb2mg==\n-----END PRIVATE KEY-----"
	raw := "-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgarFDcGR3uSB7wUlX\nTa9n803n730lq2MgbXdxUnxt9XehRANCAARWWFY5d1RmSuHiezj1I1XU6OVgWiZ/\niM1jJS9HhcWJ/I0YcmV8LMeoqduFAyXsaxjJlOcASv5AU6FrX7t0Fvaa\n-----END PRIVATE KEY-----"
	block, err := pem.Decode([]byte(raw))
	if err != nil {
		fmt.Println(err)
	}

	//if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
	//	fmt.Println(key)
	//}

	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		switch key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return
		default:
			fmt.Println("hello")
		}
	}

	var publickeyPEM []byte

	if key, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
		fmt.Println(key.PublicKey)
		publicKeyDer, err := x509.MarshalPKIXPublicKey(&(key.PublicKey))
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(base64.StdEncoding.EncodeToString(publicKeyDer))
		publickeyPEM = pem.EncodeToMemory(
			&pem.Block{
				Type:  "PUBLIC KEY",
				Bytes: publicKeyDer,
			},
		)

		fmt.Println(string(publickeyPEM))
	}
	//decodeString, err2 := base64.StdEncoding.DecodeString(string(publickeyPEM))
	decodeString, err2 := base64.StdEncoding.DecodeString("MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVlhWOXdUZkrh4ns49SNV1OjlYFomf4jNYyUvR4XFifyNGHJlfCzHqKnbhQMl7GsYyZTnAEr+QFOha1+7dBb2mg==")
	if err2 != nil {
		fmt.Println(err2.Error())

	}

	bytes, err2 := hex.DecodeString(string(decodeString))
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(string(bytes))

	if key, err := sm2.ParsePKCS8UnecryptedPrivateKey(block.Bytes); err == nil {
		fmt.Println(key)
	} else {
		fmt.Printf("error!!!!! %s", err.Error())
	}

}
func TestParsePublicKey(t *testing.T) {
	//raw:="-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEVlhWOXdUZkrh4ns49SNV1OjlYFomf4jNYyUvR4XFifyNGHJlfCzHqKnbhQMl7GsYyZTnAEr+QFOha1+7dBb2mg==\n-----END PUBLIC KEY-----"
	//raw:="-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAE2D4CFe/FkmwGW5dAsiodfmfCsEjlJjcB\nO0/mVgQ0ctKO5/h3eWq+lIgKiqJ5p5skuuDBr5ZbWCNpSLvgK+ODMg==\n-----END PUBLIC KEY-----"
	raw := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVlhWOXdUZkrh4ns49SNV1OjlYFom\nf4jNYyUvR4XFifyNGHJlfCzHqKnbhQMl7GsYyZTnAEr+QFOha1+7dBb2mg==\n-----END PUBLIC KEY-----"

	block, err := pem.Decode([]byte(raw))
	if err != nil {
		fmt.Println(err)
	}
	key, err2 := x509.ParsePKIXPublicKey(block.Bytes)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(key)

}
func TestParseCertificateRequest(t *testing.T) {
	str := "-----BEGIN CERTIFICATE REQUEST-----\nMIHJMHICAQAwEDEOMAwGA1UEAwwFYWRtaW4wWTATBgcqhkjOPQIBBggqhkjOPQMB\nBwNCAATjIeVS7uhN0qsKebneNUe5cYTwl2XWJ5BGczOQzObpwvbJFUjYSlja6iHk\n3bQaZvkcno+PQgSM6TBiES22MRY0oAAwCgYIKoZIzj0EAwIDRwAwRAIgHMVKpUDO\n93JjIet9Jt0oyLo+tAQwxr0HAF+iNaf+pEQCIFUl0k9xEWTVU8LCCVSRNgZQKyk7\nrtQEMxOBQYdMA9vX\n-----END CERTIFICATE REQUEST-----"
	decode, _ := pem.Decode([]byte(str))

	request, err := x509.ParseCertificateRequest(decode.Bytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(request)
}

func TestParseCertificate(t *testing.T) {
	path := "/Users/fengxiaoxiao/work/go/src/github.com/hyperledger/fabric-samples/asset-transfer-basic/application-typescript/src/mywallet/cert.pem"
	cabytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	caCert, err := sm2.ReadCertificateFromMem(cabytes)
	fmt.Println(caCert)

}

func TestTime(t *testing.T) {
	str := "2022-01-31T16:00:00.000Z"

	parse, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err.Error())
	}
	unix := parse.Unix()
	fmt.Println(unix)
}
func TestVerifyCert(t *testing.T) {
	caCert := "-----BEGIN CERTIFICATE-----\nMIICGzCCAcOgAwIBAgIQMn6NESWuiMxcz8Ex9cJAfDAKBggqhkjOPQQDAjBZMQsw\nCQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpSmluZzEQMA4GA1UEBxMHQmVpSmluZzEP\nMA0GA1UEChMGYWEubnBjMRUwEwYDVQQDEwx0bHNjYS5hYS5ucGMwHhcNMjIwMzE4\nMDk0MzA1WhcNMzIwMzE1MDk0MzA1WjBZMQswCQYDVQQGEwJDTjEQMA4GA1UECBMH\nQmVpSmluZzEQMA4GA1UEBxMHQmVpSmluZzEPMA0GA1UEChMGYWEubnBjMRUwEwYD\nVQQDEwx0bHNjYS5hYS5ucGMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQOnwGl\n88cePpsf/xC+BlNx/jRCYJK1ff6q2l1ptaSqoNqWJAorlqxgol3kjSfqSqn1pGah\nAowf78Iroe1U1jbno20wazAOBgNVHQ8BAf8EBAMCAaYwHQYDVR0lBBYwFAYIKwYB\nBQUHAwIGCCsGAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wKQYDVR0OBCIEIIHKv8te\nj81rkUsX8s1Dt/LU5JxLBc+XEpAG9ftxa5fvMAoGCCqGSM49BAMCA0YAMEMCIEg9\nrP+G/zNZX8EoHwwVTnc5iqfzzxMw8px0kvVGIRY6Ah9vi+hmmv2WFaTUSEwarRre\nGELLf5paffaSqgUAG7qo\n-----END CERTIFICATE-----"
	clientCert := "-----BEGIN CERTIFICATE-----\nMIICKzCCAdKgAwIBAgIQCwIdt/eVOikztxxu6i+lUDAKBggqhkjOPQQDAjBZMQsw\nCQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpSmluZzEQMA4GA1UEBxMHQmVpSmluZzEP\nMA0GA1UEChMGYWEubnBjMRUwEwYDVQQDEwx0bHNjYS5hYS5ucGMwHhcNMjIwMzE4\nMDk0MzA1WhcNMzIwMzE1MDk0MzA1WjBLMQswCQYDVQQGEwJDTjEQMA4GA1UECBMH\nQmVpSmluZzEQMA4GA1UEBxMHQmVpSmluZzEYMBYGA1UEAxMPb3JkZXJlcjEtbnBj\nLmFhMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4HpFPynPDMnuXi2u/G46F/jH\nskD1hlrinrGq8NguF69M9+vHN9iLMMumgtoDkIFN2YhIPjU0+8yZABsnTEgpSaOB\niTCBhjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUF\nBwMCMAwGA1UdEwEB/wQCMAAwKwYDVR0jBCQwIoAggcq/y16PzWuRSxfyzUO38tTk\nnEsFz5cSkAb1+3Frl+8wGgYDVR0RBBMwEYIPb3JkZXJlcjEtbnBjLmFhMAoGCCqG\nSM49BAMCA0cAMEQCIChCIGkX9z46Ss7Mx5fOAsb+Z/Haz4KFu790gndH/0mwAiBC\nK7oV9DHLzcsdSKrAIIs+XZWFadEbhVu9oGpgnQdgqg==\n-----END CERTIFICATE-----"
	cablock, _ := pem.Decode([]byte(caCert))
	if cablock == nil {
		logger.Error(GetErrorStackf(nil, "获取 msp 证书对象失败： mspCert = %s", caCert))
	}

	cacertificate, err := x509.ParseCertificate(cablock.Bytes)
	if err != nil {
		logger.Error(GetErrorStackf(nil, "获取 msp 证书对象失败： mspCert = %s", caCert))
	}

	clientblock, _ := pem.Decode([]byte(clientCert))
	if cablock == nil {
		logger.Error(GetErrorStackf(nil, "获取 msp 证书对象失败： mspCert = %s", clientCert))
	}

	clientcertificate, err := x509.ParseCertificate(clientblock.Bytes)
	if err != nil {
		logger.Error(GetErrorStackf(nil, "获取 msp 证书对象失败： mspCert = %s", clientCert))
	}

	sha256Hash.Write(clientcertificate.RawTBSCertificate)
	sum := sha256Hash.Sum(nil)

	fmt.Println(base64.StdEncoding.EncodeToString(sum))
	fmt.Println(base64.StdEncoding.EncodeToString(clientcertificate.Signature))

	//asn1 := ecdsa.VerifyASN1(cacertificate.PublicKey.(*ecdsa.PublicKey), sum, clientcertificate.Signature)
	//fmt.Println(asn1)
	signature := clientcertificate.Signature
	ecdsaSig := new(ecdsaSignature)
	if rest, err := asn1.Unmarshal(signature, ecdsaSig); err != nil {
		logger.Errorf(err.Error())
	} else if len(rest) != 0 {
		logger.Error("x509: trailing data after ECDSA signature")
	}
	if ecdsaSig.R.Sign() <= 0 || ecdsaSig.S.Sign() <= 0 {
		logger.Error("x509: ECDSA signature contained zero or negative values")
	}
	fmt.Println(ecdsa.Verify(cacertificate.PublicKey.(*ecdsa.PublicKey), sum, ecdsaSig.R, ecdsaSig.S))
}

type dsaSignature struct {
	R, S *big.Int
}

type ecdsaSignature dsaSignature
