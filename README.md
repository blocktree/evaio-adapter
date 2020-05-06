# evaio-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建EVA.ini文件，编辑如下内容：

```ini
# transaction type
txType = "cosmos-sdk/StdTx"
# message type
msgSend = "cosmos-sdk/MsgSend"
msgVote = "cosmos-sdk/MsgVote"
msgDelegate = "cosmos-sdk/MsgDelegate"
# message choose 1-send  2-vote  3-delegate
msgType = 1


# mainnet rest api url
mainnetRestAPI = "https://stargate.evaio.net"
# mainnet node api url
mainnetNodeAPI = ""
# chain id
mainnetChainID = "evaio"
# mainnet denom
mainnetDenom = "neva"

# testnet rest api url
testnetRestAPI = "http://47.91.232.118:8888"
# testnet node api url
testnetNodeAPI = "http://192.168.27.124:20041"
# chain id
testnetChainID = "evaio-dev"
# testnet denom
testnetDenom = "neva"

# Is network test?
isTestNet = false

# scan mempool or not
isScanMemPool = false

# pay fee or not
payFee = true
# minimum fee to pay in muon/uatom(1 mon = 1000000muon , 1 atom = 1000000uatom)
minFee = 10000000000
# standed gas
stdGas = 200000

# Cache data file directory, default = "", current directory: ./data
dataDir = ""
```
