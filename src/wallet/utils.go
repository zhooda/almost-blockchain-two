package wallet

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mr-tron/base58"
)

// Base58Encode encodes hash using base58 algorithm
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

// Base58Decode decodes hash using base58 algorithm
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

// haha this is not a public function, go compiler cant yell
// at me for having improper comments >:)
func getWalletFile() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	tmpPath := filepath.Join(exPath, "tmp")
	walletPath := filepath.Join(tmpPath, "wallets.data")

	_ = os.MkdirAll(tmpPath, os.ModePerm)

	return walletPath
}
