package blockchain

import (
	"bytes"

	"github.com/zhooda/almost-blockchain-two/wallet"
)

// TxOutput contains value in tokens
// and key needed to unlock tokens
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// TxInput is a reference to a previous
// output. It references the Transaction
// that contains the output, the index where
// the output is, and the signature
type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

// NewTXOutput creates a new txoutput
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// UsesKey compares txinput public key with
// argument public key
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// Lock will lock a txoutput
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks to see if output has been locked with key
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
