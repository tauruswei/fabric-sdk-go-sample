package main

import (
	"fabric-go-sdk-sample/log"
	"fabric-go-sdk-sample/router"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"strconv"
)

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2023/4/26 16:41
 */

const config_yaml = "./config.yaml"

func main() {

	//service.NewService(config_yaml)

	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	route := router.CreateRouter()
	portString := "8081"
	port, err := strconv.Atoi(portString)
	if err != nil {
		panic("parse server port:" + portString + " err: " + err.Error())
	}
	route.Use(TlsHandler(port))
	// todo listenAddress   tls.server.cert   tls.server.key

	route.Run(":" + "8081")

}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			return
		}

		c.Next()
	}
}
