Fabric sdk go sample
==========
基于 fabric-samples/first-network启动的网络,通过使用 fabric-sdk-go 来操作链码的例子
#
`environment:`      
`golang v1.13.4`  `fabric v1.4.2`  `fabric-samples v1.4.2`
#
`images:`  
 `hyperledger/fabric-peer:1.4.2`  `hyperledger/fabric-orderer:1.4.2`  `hyperledger/fabric-tools:1.4.2`  `hyperledger
 /fabric-ccenv:1.4.2`  `hyperledger/fabric-baseos:amd64-0.4.15`
#


目录:

- app: main.go，测试的主程序
- binaries: 二进制文件夹，需要将所有的二进制文件拷贝到 ***/fabric-samples/bin/ 目录下, 并在$PATH 变量中添加 ***/fabric-samples/bin/
- chaincode: 链码
- cli: 链码调用代码封装
- config: fabric sdk 与区块链交互配置


## Quick start

1. 启动fabric网络    
    ```
    进入 fabric-samples/first-network 
    ./byfn.sh up
    ```


2. 配置`config.yaml`
    ```
    path: ***/fabric-samples/first-network/crypto-config
    *** 改为自己的相关目录
    ```

3. 依赖包的安装
    ```
    cd fabric-sdk-go-sample
    go mod download
    ```

 
4. 运行客户端程序`go run main.go`

    ```bash
    [root@jc chaincode]# cd app
    [root@jc chaincode]# go run main.go
    2020/03/26 21:20:27 Initialized fabric sdk
    2020/03/26 21:20:27 Initialized resource client
    2020/03/26 21:20:27 Initialized channel client
    2020/03/26 21:20:27 =================== Phase 1 begin ===================
    2020/03/26 21:20:27 Install  response status: 200
    2020/03/26 21:20:27 Chaincode has been installed on org1's peers
    2020/03/26 21:20:47 Instantitate chaincode tx: 7bc22a9e65766b4c38451f4f1271c5c961fd1d7a5880da1cd2141d322ce2139c
    2020/03/26 21:20:47 Chaincode has been instantiated
    2020/03/26 21:20:49 Invoke chaincode response:
    id: aecd053b42e591f8e493236b281e8379edc156373004c6b98aefdd5191a0be2b
    validate: VALID
    chaincode status: 200

    2020/03/26 21:20:49 Invoke chaincode success
    2020/03/26 21:20:49 Query chaincode tx response:
    tx: 2a1f7f060f19cb2a6f9a5ec30180d57ba820b00c88ce794eb6a3984ed629e083
    result: 90

    2020/03/26 21:20:49 Query chaincode success on peer0.org1
    2020/03/26 21:20:49 =================== Phase 1 end ===================
    ```