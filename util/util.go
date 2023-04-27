package util

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
	rand1 "math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	// Curve 椭圆曲线
	Curve = elliptic.P256()
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/21 下午4:39
 */
///**
//  签名
//*/
//func Sign(data []byte) ([]byte, error) {
//	hash, err := global.Verifier.CSP.Hash(data, &bccsp.SHA256Opts{})
//	if err != nil {
//		return nil, errors.WithMessagef(err, "计算哈希失败, data = %s ", base64.StdEncoding.EncodeToString(data))
//	}
//	signatrure, err := global.Verifier.Sign(global.Cert.SubjectKeyId, hash)
//	if err != nil {
//		return nil, err
//	}
//	return signatrure, nil
//}
//
///**
//  验签
//*/
//func Verify(sig, data, certBytes []byte) (bool, error) {
//	hash, err := global.Verifier.CSP.Hash(data, &bccsp.SHA256Opts{})
//	if err != nil {
//		return false, errors.WithMessagef(err, "计算哈希失败, data = %s ", base64.StdEncoding.EncodeToString(data))
//	}
//	cert, err := sm2.ReadCertificateFromMem(certBytes)
//	if err != nil {
//		return false, errors.WithMessagef(err, "解析签名证书失败, cert = %s", string(certBytes))
//	}
//	result, err := global.Verifier.Verify(cert.SubjectKeyId, sig, hash)
//	if err != nil {
//		return false, err
//	}
//	return result, nil
//}

func MakeTempdir() string {
	dir := os.TempDir()
	tempDir := filepath.Join(dir, "CaTemp")
	intermediateDir := ""
	intermediateDir = RandStringInt()
	return filepath.Join(tempDir, intermediateDir)
}

// 产生随机数
func RandStringInt() string {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	return serialNumber.String()
}
func RandCodes() string {
	rnd := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
func Uint64Rand() uint64 {
	return uint64(rand1.Uint32())<<32 + uint64(rand1.Uint32())
}
func ParseTime(timeString string) (string, error) {
	if strings.Contains(timeString, "Z") {
		parse, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			return "", err
		}
		unix := parse.Format("2006-01-02 15:04:05")
		return fmt.Sprintf("%s", unix), nil
	}
	return timeString, nil

}
