package core

type TXInput struct {
	Txid      []byte // 存放这个输入端所依赖的交易 ID
	Vout      int    // 存放交易的输出索引
	ScriptSig string // 用于解锁输出端
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
