package v1

import (
	"fabric-go-sdk-sample/model"
	"fabric-go-sdk-sample/result"
	"fabric-go-sdk-sample/service"
	"fabric-go-sdk-sample/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Invoke(c *gin.Context) {
	g := result.Gin{C: c}

	request := model.InvokeRequest{}

	err := util.Validator(c, &request, g)
	if err != nil {
		g.Error(result.PARAMETER_VALID_ERROR.FillArgs(err.Error()))
		return
	}

	a := []string{"set", request.Token, request.Token}
	ret, err := service.App.Set(a)
	if err != nil {
		g.Error(result.SERVER_ERROR.FillArgs(err.Error()))
		return
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	g.Success("")
}

func Query(c *gin.Context) {
	g := result.Gin{C: c}

	request := model.InvokeRequest{}

	err := util.Validator(c, &request, g)
	if err != nil {
		g.Error(result.PARAMETER_VALID_ERROR.FillArgs(err.Error()))
		return
	}

	a := []string{"get", request.Token}
	response, err := service.App.Get(a)
	if err != nil {
		g.Error(result.SERVER_ERROR.FillArgs(err.Error()))
		return
	}
	fmt.Println("<--- 查询信息　--->：", response)

	g.Success("")
}
