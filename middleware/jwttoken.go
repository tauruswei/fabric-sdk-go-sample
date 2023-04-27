package middleware

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2021/12/8 2:19 下午
 */
import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	Expiration = 1
)

// 证书签名密钥
var jwtKey = []byte("abc")

// 签名需要传递的参数
type HmacUser struct {
	Id       string `json:"id"`
	NickName string `json:"nickName"`
}

type MyClaims struct {
	Id       string `json:"id"`
	NickName string `json:"nickName"`
	jwt.StandardClaims
}

// 登录的参数
type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 定义生成token的方法
func GenerateToken(u HmacUser) (string, error) {
	// 定义过期时间,7天后过期
	expirationTime := time.Now().Add(Expiration * time.Hour)
	claims := &MyClaims{
		Id:       u.Id,
		NickName: u.NickName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发布时间
			Subject:   "token",               // 主题
			Issuer:    "水痕",                  // 发布者
		},
	}
	// 注意单词别写错了
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 定义解析token的方法
func parseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

// 定义中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestURL := c.Request.URL.String()
		switch {
		case strings.Contains(requestURL, "genMnemonics"):
		//case strings.Contains(requestURL, "genRandomCodes"):
		case strings.Contains(requestURL, "login"):
		//case strings.Contains(requestURL, "registerUser"):
		case strings.Contains(requestURL, "signCertificate"):
		case strings.Contains(requestURL, "nftList"):
		case strings.Contains(requestURL, "downloadConfigFile"):
		case strings.Contains(requestURL, "queryNFTListByHot"):
		case strings.Contains(requestURL, "queryNFTListLatest"):
		case strings.Contains(requestURL, "queryNFTTokenHistory"):
		case strings.Contains(requestURL, "queryNFTDetail"):
		case strings.Contains(requestURL, "getReleaseList"):
		case strings.Contains(requestURL, "getContactList"):
		case strings.Contains(requestURL, "saveContact"):
		case strings.Contains(requestURL, "userExist"):
			// case strings.Contains(requestURL, "getMessage"):

			//case strings.Contains(requestURL, "getOssToken"):
			//case strings.Contains(requestURL, "createNft721"):
			//case strings.Contains(requestURL, "nft"):
			c.Next()
			return
		default:
			// 从请求头中获取token
			tokeString := c.GetHeader("Authorization")
			//fmt.Println(tokeString, "当前token")
			if tokeString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "必须传递token",
				})
				c.Abort()
				return
			}
			token, claims, err := parseToken(tokeString)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  fmt.Sprintf("token解析错误，err = %s", err.Error()),
				})
				c.Abort()
				return
			}
			// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
			c.Set("id", claims.Id)
			c.Set("nickName", claims.NickName)
			c.Next()

		}
	}
}
