module fabric-go-sdk-sample

go 1.16

replace (
	github.com/hyperledger/fabric-sdk-go v1.0.0 => ./third_party/fabric-sdk-go
	github.com/tjfoc/gmsm => ./third_party/gmsm
)

require (
	github.com/hyperledger/fabric-protos-go v0.0.0-20210318103044-13fdee960194
	github.com/hyperledger/fabric-sdk-go v1.0.0
)
