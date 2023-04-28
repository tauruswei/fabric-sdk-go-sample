# fabric-go-sdk-sample

本项目基于hyperledger fabric 2.x网络

## 基本流程

### 编译国密镜像

```
git clone https://github.com/tauruswei/fabric
cd fabric
# fabric 版本 2.2.4 (以 arm64 版本为例，官方的 amd64 镜像也可以)
git checkout -b 2.2.4-aarch64 2.2.4-aarch64
go env -w GO111MODULE=on
go mod tidy
go mod vendor
make peer-docker-clean
make peer-docker
make orderer-docker-clean
make orderer-docker
make ccenv-docker-clean
make ccenv-docker
make baseos-docker-clean
make baseos-docker
```

### 拉取项目

```
git clone https://github.com/tauruswei/fabric-sdk-go-sample
cd ./fabric-go-sdk-sample/fixtures/ 
git checkout -b invokeAndQuery-server invokeAndQuery-server
```


### 启动项目

```
# 修改 config.yaml 文件：节点 ip 和 crypto-config path
cd .. && go build && ./fabric-go-sdk-sample
```
```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /fabric/init              --> fabric-go-sdk-sample/router/api/v1.Init (4 handlers)
[GIN-debug] POST   /fabric/invoke            --> fabric-go-sdk-sample/router/api/v1.Invoke (4 handlers)
[GIN-debug] POST   /fabric/query             --> fabric-go-sdk-sample/router/api/v1.Query (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8081
```

### 初始化项目

```
http://127.0.0.1:8081/fabric/init
```
```
>> 开始创建通道......
>>>> 使用每个org的管理员身份更新锚节点配置...
>>>> 使用每个org的管理员身份更新锚节点配置完成
>> 创建通道成功
>> 加入通道......
>> 加入通道成功
>> 开始打包链码......
>> 打包链码成功
>> 开始安装链码......
>> 安装链码成功
>> 组织认可智能合约定义......
>>> chaincode approved by Org1 peers:
        grpcs://192.168.2.150:7051
>>> chaincode approved by Org2 peers:
        grpcs://192.168.2.150:9051
>> 组织认可智能合约定义完成
>> 检查智能合约是否就绪......
LifecycleCheckCCCommitReadiness cc = samplecc, = {map[Org1MSP:true Org2MSP:true]}
LifecycleCheckCCCommitReadiness cc = samplecc, = {map[Org1MSP:true Org2MSP:true]}
>> 智能合约已经就绪
>> 提交智能合约定义......
>> 智能合约定义提交完成
>> 调用智能合约初始化方法......
>> 完成智能合约初始化
>> 通过链码外部服务设置链码状态......
```


### 调用 invoke 接口

```
http://127.0.0.1:8081/fabric/invoke
# 消息体
{  
  "token": "1" 
} 

```
```
<--- 添加信息　--->： c4f3607dc0e1ab9601a6e7df74f8eb5a735f529aa4c5056be19c0bfa81b5c22f
[GIN] 2023/04/28 - 14:32:21 | 200 |  2.055227625s |   192.168.2.150 | POST     "/fabric/invoke"
```

### 调用 query 接口
```
http://127.0.0.1:8081/fabric/query
# 消息体
{  
  "token": "1" 
} 
```
```
<--- 查询信息　--->： 1
[GIN] 2023/04/28 - 14:32:25 | 200 |   17.773166ms |   192.168.2.150 | POST     "/fabric/query"
```

### 清理项目

```
cd ./fabric-go-sdk-sample/fixtures/ && bash stop.sh
```
