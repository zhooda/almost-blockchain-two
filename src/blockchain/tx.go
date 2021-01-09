package blockchain

// TxOutput contains value in tokens
// and key needed to unlock tokens
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput is a reference to a previous
// output. It references the Transaction
// that contains the output, the index where
// the output is, and the signature
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock verifies if a recipient can unlock a TxInput
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked verifies if a TxOutput can be unlocked by a recipient
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
