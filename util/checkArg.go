package util

import (
	"encoding/json"
	"fabric-go-sdk-sample/config"
	"fabric-go-sdk-sample/log"
	"fabric-go-sdk-sample/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"reflect"
)

func CheckCCArg(ccName string, eventFilter string,
	callbackUrl string, serisalNumber string) string {

	if len(ccName) == 0 {
		return "智能合约名称为空!"
	} else if len(ccName) > 64 {
		return "智能合约名称超长!"
	}

	if len(eventFilter) == 0 {
		return "事件过滤器为空!"
	} else if len(eventFilter) > 64 {
		return "事件过滤器名称超长!"
	}

	//if len(callbackUrl) == 0 || callbackUrl == "" {
	//	return "回调路径为空!"
	//}

	//if len(serisalNumber) == 0 {
	//	return "业务流水号为空!"
	//} else if len(serisalNumber) > 32 {
	//	return "业务流水号超长！"
	//}
	return ""
}

func CheckUnCCArg(ccName string, eventFilter string,
	serisalNumber string) string {

	if len(ccName) == 0 {
		return "智能合约名称为空!"
	} else if len(ccName) > 64 {
		return "智能合约名称超长!"
	}

	if len(eventFilter) == 0 {
		return "事件过滤器为空"
	} else if len(eventFilter) > 64 {
		return "事件过滤器名称超长!"
	}

	if len(serisalNumber) == 0 {
		return "业务流水号为空"
	} else if len(serisalNumber) > 32 {
		return "业务流水号超长！"
	}

	return ""
}

func CheckUnBlcArg(serisalNumber string) string {
	if len(serisalNumber) == 0 {
		return "业务流水号为空"
	} else if len(serisalNumber) > 32 {
		return "业务流水号超长！"
	}

	return ""
}

func Validator(c *gin.Context, v interface{}, g result.Gin) error {
	err := c.ShouldBind(v)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			log.Error(GetErrorStack(err, ""))
			return err
		}
		for _, validationErr := range errs {
			fieldName := validationErr.Field() //获取是哪个字段不符合格式
			typeOfField := reflect.TypeOf(v)
			if typeOfField.Kind() == reflect.Ptr {
				typeOfField = typeOfField.Elem()
			}
			field, ok := typeOfField.FieldByName(fieldName) //通过反射获取filed
			if ok {
				errorInfo := field.Tag.Get("reg_error_info") //获取field对应的reg_error_info tag值
				if errorInfo != "" {
					//return fieldName + ":" + errorInfo           //返回错误
					return errors.New(fmt.Sprintf(fieldName + ":" + errorInfo))
				}

			}
		}
		errormsg, err1 := json.Marshal(errs.Translate(config.Trans))
		if err1 != nil {
			log.Error(GetErrorStack(err1, ""))
			return err1
		}
		log.Error(GetErrorStack(errors.Errorf("参数验证异常: %s", string(errormsg)), ""))
		return errors.Errorf("参数验证异常: %s", string(errormsg))
	}
	return nil
}
