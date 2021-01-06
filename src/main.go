package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain structure type
type BlockChain struct {
	blocks []*Block
}

// Block structure type
type Block struct {
	Hash     []byte // hash of current block
	Data     []byte // data inside the block
	PrevHash []byte // hash of previous block
}

// DeriveHash creates block hash
// is a method for Block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock returns block
// data: string of data to be stored in block
// prevHash: byte array of previous blocks hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock adds new block to the chain
// is a method for BlockChain
// data: string of data to be stored in block
func (chain *BlockChain) AddBlock(data string) {
	// get previous block in the chain
	prevBlock := chain.blocks[len(chain.blocks)-1]

	// create new block
	new := CreateBlock(data, prevBlock.Hash)

	// add new block to the chain
	chain.blocks = append(chain.blocks, new)
}

// Genesis returns genesis block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain returns BlockChain with genesis Block
// MARK: Rename this to NewBlockChain later to comply with Go style guides
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for i, block := range chain.blocks {
		fmt.Printf("Block index: %d\n", i)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
