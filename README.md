Fabric sdk go sample
==========
基于 fabric-samples/test-network启动的网络,通过使用 fabric-sdk-go 来操作链码的例子,在fabric2.2.2 版本上，调试 chaincode external service
#
`environment:`      
`golang v1.14.12`  `fabric v2.2.2`  `fabric-samples v2.2.2`
#
`images:`  
 `hyperledger/fabric-peer:2.2.2`  `hyperledger/fabric-orderer:2.2.2`  `hyperledger/fabric-tools:2.2.2`  `hyperledger
 /fabric-ccenv:2.2.2` 
#


目录:

- app: main.go，测试的主程序
- chaincode: 链码
- cli: 链码调用代码封装
- config: fabric sdk 与区块链交互配置


## Quick start

1. 启动fabric网络    

教程：https://github.com/tauruswei/fabric-samples/tree/release2.2.2/asset-transfer-basic/chaincode-external


2. 配置`config.yaml`
    ```
    path: ***/fabric-samples/first-network/crypto-config
    *** 改为自己的相关目录
    ```

3. 修改 admin 证书名称
    ```
    mv fabric-samples\test-network\organizations\peerOrganizations\org1.example.com\users\Admin@org1.example.com\msp\signcerts\cert.pem fabric-samples\test-network\organizations\peerOrganizations\org1.example.com\users\Admin@org1.example.com\msp\signcerts\Admin@org1.example.com-cert.pem
    ```

 
4. 运行客户端程序`go run main.go`
