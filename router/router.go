package router

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:18
 */

import (
	"fabric-go-sdk-sample/middleware"
	v1 "fabric-go-sdk-sample/router/api/v1"

	"github.com/gin-gonic/gin"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateRouter 生成路由
func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors()) //开启中间件 允许使用跨域请求
	// swagger
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// invoke
	farbicGroup := router.Group("/fabric")
	farbicGroup.POST("/invoke", v1.Invoke) // invoke
	farbicGroup.POST("/query", v1.Query)   // query

	return router
}
