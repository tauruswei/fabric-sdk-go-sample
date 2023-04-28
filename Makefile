BINARY=fabric-sdk-go-sample-armv7

# 编译 适用于 armv7 架构的可执行文件
build-arm:
	docker run --rm -v $(PWD):/output/fabric-sdk-go-sample --platform=linux/arm/v7 golang:1.16 sh -c "go env -w GO111MODULE=on && cd /output/fabric-sdk-go-sample &&  go build -o ${BINARY} main.go"