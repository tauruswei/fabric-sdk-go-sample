package util

import "fabric-go-sdk-sample/model"

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2022/1/2 5:10 下午
 */
/*
 * @Desc: 设置 client 的配置
 * @Param:
 * @Return:
 */
func SetClient(config *model.ConnectionConfig, orgName, channelName, contractName, timeout string) model.Client {
	client := model.Client{
		Organization: orgName,
		ChannelName:  channelName,
		ContractName: contractName,
	}
	client.Connection.Timeout.Peer.Endorder = timeout
	client.Connection.Timeout.Orderer = timeout
	return client

}

/*
 * @Desc: 设置 单个 channel 的 peers 和 orderers
 * @Param:
 * @Return:
 */
func SetChannel(peerNames, ordererNames []string) model.Channel {
	channel := model.Channel{
		Orderers: ordererNames,
	}
	peers := make(map[string]model.Peer)
	for _, peerName := range peerNames {
		peers[peerName] = model.Peer{
			EndorsingPeer:  true,
			ChaincodeQuery: true,
			LedgerQuery:    true,
			EventSource:    true,
		}
	}
	channel.Peers = peers
	return channel
}

/*
 * @Desc: 设置单个 organization 下的 mspid 和 peers
 * @Param:
 * @Return:
 */
func SetOrganization(mspId string, peerNames []string) model.Organization {
	return model.Organization{
		MspId: mspId,
		Peers: peerNames,
	}
}

/*
 * @Desc: 设置 orderer 或者 peer 节点
 * @Param:
 * @Return:
 */
func SetNode(nodeName, url, mspId, certPem, sslTargetNameOverride, hostnameOverride string) model.Node {
	node := model.Node{
		Url:   url,
		Mspid: mspId,
	}
	node.TlsCACerts.Pem = certPem
	node.GrpcOptions.SslTargetNameOverride = nodeName
	node.GrpcOptions.HostnameOverride = nodeName
	return node
}
