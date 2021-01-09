package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// Block contains sha256 sum of self, string data,
// and sha256 sum of previous block
type Block struct {
	Hash         []byte         // hash of current block
	Transactions []*Transaction // array of transactions
	PrevHash     []byte         // hash of previous block
	Nonce        int            // nonce for PoW
}

// HashTransactions returns a sha256 hash of all
// transactions combined
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// CreateBlock returns block
// data: string of data to be stored in block
// prevHash: byte array of previous blocks hash
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// Genesis returns genesis block in the chain
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// Serialize will serialize a block for storage in
// the badger database
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// Deserialize will deserialize a byte slice from
// the database into a block struct
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Handle will panic if an error is not nil
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
