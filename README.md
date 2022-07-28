# fabric-go-sdk-sample

本项目基于hyperledger fabric 2.x网络

## 基本流程

### 编译国密镜像

```
git clone https://github.com/tauruswei/fabric
cd fabric
git checkout -b 2.2.4-gm-withour-tls-aarch64 2.2.4-gm-withour-tls-aarch64
go env -w GO11MODULE=on
go mod tidy
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
git checkout -b 2.2.4-1.0.0-btea3-gm-without-tls-aarch64 2.2.4-1.0.0-btea3-gm-without-tls-aarch64
```

### 启动节点

```
docker-compose up -d
```

### 启动项目

```
cd .. && go build && ./fabric-go-sdk-sample
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
	peer0.org1.example.com:7051
>>> chaincode approved by Org2 peers:
        grpcs://localhost:9051
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
>> 设置链码状态完成
<--- 添加信息　--->： 18c0c86ce029d7de04461484976c5151992864b52ca28905d0ccf911443fdfcb
<--- 查询信息　--->： 123
```

