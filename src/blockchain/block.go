package blockchain

// BlockChain contains slice of Block
type BlockChain struct {
	Blocks []*Block
}

// Block contains sha256 sum of self, string data,
// and sha256 sum of previous block
type Block struct {
	Hash     []byte // hash of current block
	Data     []byte // data inside the block
	PrevHash []byte // hash of previous block
	Nonce    int    // nonce for PoW
}

// CreateBlock returns block
// data: string of data to be stored in block
// prevHash: byte array of previous blocks hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// AddBlock adds new block to the chain
// is a method for BlockChain
// data: string of data to be stored in block
func (chain *BlockChain) AddBlock(data string) {
	// get previous block in the chain
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	// create new block
	new := CreateBlock(data, prevBlock.Hash)

	// add new block to the chain
	chain.Blocks = append(chain.Blocks, new)
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
