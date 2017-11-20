package core

import "bytes"

type TXInput struct {
	Txid      []byte // 存放这个输入端所依赖的交易 ID
	Vout      int    // 存放交易的输出索引
	Signature []byte // 用于解锁输出端
	PubKey    []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
