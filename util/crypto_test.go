package util

import (
	"fmt"
	"testing"
)

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2022/1/15 2:09 下午
 */
func TestGenKeyPem(t *testing.T) {
	tlsStr := "dGlueSBzZWVkIGNvdHRvbiBzaW5nIHN3YXAgYmVuZWZpdCBhbHJlYWR5IGJveCBtZW51IHJ1bGUgd2FnZSBwZW4gc29vbiBzaG9jayBzb25nIGNhcmdvIGRpYWdyYW0gcmFtcA=="
	//tlsKeyBytes, err := base64.StdEncoding.DecodeString(tlsStr)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	pem, err := GenKeyPem(tlsStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(pem))
	mspKeyPem := pem

	mspKey, _, err := ImportPrivateKey(0, string(mspKeyPem), "SW", "ECC", "test")
	fmt.Println(mspKey)
}
