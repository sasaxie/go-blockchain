# go-blockchain

# Installation

```
> go get github.com/sasaxie/go-blockchain
> cd $GOPATH/src/github.com/sasaxie/go-blockchain/
> go build
> ./go-blockchain [command] -[option]
```

# Commands

**Create wallet:**

```
> ./go-blockchain createwallet
> Your new address: 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119
> ./go-blockchain createwallet
> Your new address: 1546AtPEFt5uGyvhbRDLUoqLTy5zfUZhsY
```

**Create blockchain:**

```
> ./go-blockchain createblockchain -address 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119
> 000000a69db5326fa6037956b4cfef0ecdc029b3de3ddad59bf0940944653510
> 
> Done!
```

**Reindex UTXO:**

```
> ./go-blockchain reindexutxo
> Done! There are 1 transactions in the UTXO set.
```

**Get balance:**

```
> ./go-blockchain getbalance -address 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119
> Balance of '1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119': 10
```

**Send coin:**

```
> ./go-blockchain send -from 1321QKXdSdYXLPZiGxC9NS6HjTuDEFR119 -to 1546AtPEFt5uGyvhbRDLUoqLTy5zfUZhsY -amount 1 -mine
> 00000013eda8154a87b0ee7ad5bf54f63cd6a0a2de89ce4bbc9f08313cf2c62a
> 
> Success!
```

**Start node:**

```
> ./go-blockchain startnode
```

**Set miner node:**

```
> ./go-blockchain startnode -miner 14kiRgXmK5tDq3MzoM33Tz3wLZSTbBK9TK
```

**Print chain:**

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

# Example

Let’s play the scenario we defined earlier.

First, set **NODE_ID** to 3000 (**export NODE_ID=3000**) in the first terminal window. I’ll use badges like **NODE 3000** or **NODE 3001** before next paragraphs, for you to know what node to perform actions on.

**NODE 3000**
Create a wallet and a new blockchain:

```
$ ./go-blockchain createblockchain -address CENTREAL_NODE
```

(I’ll use fake addresses for clarity and brevity)

After that, the blockchain will contain single genesis block. We need to save the block and use it in other nodes. Genesis blocks serve as identifiers of blockchains (in Bitcoin Core, the genesis block is hardcoded).

```
$ cp blockchain_3000.db blockchain_genesis.db 
```

**NODE 3001**
Next, open a new terminal window and set node ID to 3001. This will be a wallet node. Generate some addresses with **./go-blockchain createwallet**, we’ll call these addresses **WALLET_1**, **WALLET_2**, **WALLET_3**.

**NODE 3000**
Send some coins to the wallet addresses:

```
$ ./go-blockchain send -from CENTREAL_NODE -to WALLET_1 -amount 10 -mine
$ ./go-blockchain send -from CENTREAL_NODE -to WALLET_2 -amount 10 -mine
```

**-mine** flag means that the block will be immediately mined by the same node. We have to have this flag because initially there are no miner nodes in the network.
Start the node:

```
$ ./go-blockchain startnode
```

The node must be running until the end of the scenario.

**NODE 3001**
Start the node’s blockchain with the genesis block saved above:

```
$ cp blockchain_genesis.db blockchain_3001.db
```

Run the node:

```
$ ./go-blockchain startnode
```

It’ll download all the blocks from the central node. To check that everything’s ok, stop the node and check the balances:

```
$ ./go-blockchain getbalance -address WALLET_1
Balance of 'WALLET_1': 10

$ ./go-blockchain getbalance -address WALLET_2
Balance of 'WALLET_2': 10
```

Also, you can check the balance of the **CENTRAL_NODE** address, because the node 3001 now has its blockchain:

```
$ ./go-blockchain getbalance -address CENTRAL_NODE
Balance of 'CENTRAL_NODE': 10
```

**NODE 3002**
Open a new terminal window and set its ID to 3002, and generate a wallet. This will be a miner node. Initialize the blockchain:

```
$ cp blockchain_genesis.db blockchain_3002.db
```

And start the node:

```
$ ./go-blockchain startnode -miner MINER_WALLET
```

**NODE 3001**
Send some coins:

```
$ ./go-blockchain send -from WALLET_1 -to WALLET_3 -amount 1
$ ./go-blockchain send -from WALLET_2 -to WALLET_4 -amount 1
```

**NODE 3002**
Quickly! Switch to the miner node and see it mining a new block! Also, check the output of the central node.

**NODE 3001**
Switch to the wallet node and start it:

```
$ ./go-blockchain startnode
```

It’ll download the newly mined block!

Stop it and check balances:

```
$ ./go-blockchain getbalance -address WALLET_1
Balance of 'WALLET_1': 9

$ ./go-blockchain getbalance -address WALLET_2
Balance of 'WALLET_2': 9

$ ./go-blockchain getbalance -address WALLET_3
Balance of 'WALLET_3': 1

$ ./go-blockchain getbalance -address WALLET_4
Balance of 'WALLET_4': 1

$ ./go-blockchain getbalance -address MINER_WALLET
Balance of 'MINER_WALLET': 10
```

That’s it!
