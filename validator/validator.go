package validator

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2021/12/11 10:57 下午
 */

const (
	//密码验证选项
	PWD_OPT_Number  uint16 = 1 << iota //数字 	 0001
	PWD_OPT_Lower                      //小写 	 0010
	PWD_OPT_Upper                      //大写 	 0100
	PWD_OPT_Special                    //特殊符号 1000
)

func VerifyPwd(pwd string, options uint16, must uint16) bool {
	if pwd == "" {
		return false
	}
	var result uint16
	for _, r := range pwd {
		switch {
		case unicode.IsNumber(r):
			result = result | PWD_OPT_Number
		case unicode.IsLower(r):
			result = result | PWD_OPT_Lower
		case unicode.IsUpper(r):
			result = result | PWD_OPT_Upper
		case unicode.IsPunct(r) || unicode.IsSymbol(r): //标点符号 和 字符
			result = result | PWD_OPT_Special
		default:
			return false
		}
		// 比较结果和设置项
		// 当 cp.options&cp.result != cp.result 表示密码字符串超出 options 范围
		if options&result != result {
			return false
		}
	}

	return must&result == must
}

func Passwd(fl validator.FieldLevel) bool {
	if passwd, ok := fl.Field().Interface().(string); ok {
		pwdoptall := PWD_OPT_Lower | PWD_OPT_Number | PWD_OPT_Upper | PWD_OPT_Special
		flag := VerifyPwd(passwd, pwdoptall, pwdoptall)
		//if !flag {
		//	c.JSON(http.StatusBadRequest, model.NewErrorResponse(errors.New("密码必须有大写字母，小写字母，字符和数字")))
		//	return
		//}
		return flag
	}
	return true

}
