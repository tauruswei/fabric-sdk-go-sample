package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fabric-go-sdk-sample/global"
	logger "fabric-go-sdk-sample/log"
	"fmt"
	"github.com/pkg/errors"
	"hash"
	"io"
	"time"
)

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2021/12/14 11:45 上午
 */

/*
 * @Desc:
 * @Param:
 * @Return:
 */
func GetPolicyToken(nickName string) (string, error) {
	now := time.Now().Unix()
	expire_end := now + global.OssConfig.Expiration
	var tokenExpire = get_gmt_iso8601(expire_end)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, fmt.Sprintf("%s%s/", global.OssConfig.UploadDir, nickName))
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(global.OssConfig.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = global.OssConfig.CallbackUrl
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callback_str, err := json.Marshal(callbackParam)
	if err != nil {
		logger.Error(GetErrorStackf(err, "callback json error"))
		return "", errors.WithMessagef(err, "callback json error")
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callback_str)

	var policyToken PolicyToken
	policyToken.AccessKeyId = global.OssConfig.AccessKeyId
	policyToken.Host = global.OssConfig.Endpoint
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = fmt.Sprintf("%s%s/", global.OssConfig.UploadDir, nickName)
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)
	response, err := json.Marshal(policyToken)
	if err != nil {
		logger.Error(GetErrorStackf(err, ""))
		return "", errors.WithMessagef(err, "")
	}
	return string(response), nil
}
func get_gmt_iso8601(expire_end int64) string {
	var tokenExpire = time.Unix(expire_end, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}
type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}
type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}
