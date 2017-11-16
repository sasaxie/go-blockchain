package core

type TXOutput struct {
	Value        int    // coin
	ScriptPubKey string // 作为新的输入端时需要使用该值进行解锁
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
