module fabric-go-sdk-sample

go 1.16

replace (
	github.com/gin-gonic/gin v1.9.0 => github.com/gin-gonic/gin v1.7.7
	github.com/hyperledger/fabric-sdk-go v1.0.0 => github.com/hyperledger/fabric-sdk-go v1.0.0-beta3
	github.com/tjfoc/gmsm v1.4.1 => ./third_party/tjfoc/gmsm

)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.0
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/hyperledger/fabric v1.4.12
	github.com/hyperledger/fabric-protos-go v0.0.0-20210318103044-13fdee960194
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.6.0
	github.com/rogpeppe/go-internal v1.3.0
	github.com/spf13/viper v1.8.1
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	github.com/tjfoc/gmsm v1.4.1
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/unrolled/secure v1.13.0
	gorm.io/gorm v1.25.0
)
