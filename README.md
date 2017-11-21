# go-blockchain

First, set **NODE_ID** to 3000 (**export NODE_ID=3000**) in the first terminal window. I’ll use badges like **NODE 3000** or **NODE 3001** before next paragraphs, for you to know what node to perform actions on.

**创建钱包**

```
> ./go-blockchain createwallet
> Your new address: 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119
> ./go-blockchain createwallet
> Your new address: 1546AtPEFt5uGyvhbRDLUoqLTy5zfUZhsY
```

**创建区块链：**

```
> ./go-blockchain createblockchain -address 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119
> 000000a69db5326fa6037956b4cfef0ecdc029b3de3ddad59bf0940944653510
> 
> Done!
```

**重置 UTXO**

```
> ./go-blockchain reindexutxo
> Done! There are 1 transactions in the UTXO set.
```

**查看账号余额：**

```
> ./go-blockchain getbalance -address 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119
> Balance of '1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119': 10
```

**转账：**

```
> ./go-blockchain send -from 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119 -to 1546AtPEFt5uGyvhbRDLUoqLTy5zfUZhsY -amount 1 -mine
> 00000013eda8154a87b0ee7ad5bf54f63cd6a0a2de89ce4bbc9f08313cf2c62a
> 
> Success!
```

**启动节点**

```
> ./go-blockchain startnode
```

**设置挖矿节点**

```
> ./go-blockchain startnode -miner 14kiRgXmK5tDq3MzoM33Tz3wLZSTbBK9TK
```

**打印区块链：**

```
> ./go-blockchain printchain
> ============ Block 00000013eda8154a87b0ee7ad5bf54f63cd6a0a2de89ce4bbc9f08313cf2c62a ============
> Prev. block: 000000a69db5326fa6037956b4cfef0ecdc029b3de3ddad59bf0940944653510
> PoW: true
> 
> --- Transaction d3c67d1911c64c6287c9eafabbd3c6787a512cbf54bd44eb3dfe0edb259ef46c:
>      Input 0:
>        TXID:      a90fa883b56fad4b2dc0bb2c50a4c78430b1ac0bef5d3cd62cf836a2878e0a9a
>        Out:       0
>        Signature: cf16e093bace5b843edef1e272f9cdabd47b741265c4c3a2e651ab6f3c2a466d396621df57bd48d09bce686962e7b34ca66c83e5bb6ca4c4204661d488fe532b
>        PubKey:    a09f43b8c60eef9f511cb8ac2e79c2a6ed8cc6c8bae3ad4f71046538094ecff297fd6df23a6f3f2b6a0345594043bfe65e1e671efdf68a5e06acfc24141f65df
>      Output 0:
>        Value:  1
>        Script: 2c76112cf3bcf8514978b97055563f92fdb3b966
>      Output 1:
>        Value:  9
>        Script: 162100ce8ca041f88ad7089aa0b06a99862a2969
> 
> 
> ============ Block 000000a69db5326fa6037956b4cfef0ecdc029b3de3ddad59bf0940944653510 ============
> Prev. block: 
> PoW: true
> 
> --- Transaction a90fa883b56fad4b2dc0bb2c50a4c78430b1ac0bef5d3cd62cf836a2878e0a9a:
>      Input 0:
>        TXID:      
>        Out:       -1
>        Signature: 
>        PubKey:    5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73
>      Output 0:
>        Value:  10
>        Script: 162100ce8ca041f88ad7089aa0b06a99862a2969
```


