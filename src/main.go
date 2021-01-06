package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/zhooda/almost-blockchain-two/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")
	chain.AddBlock("Fourth block after Genesis")
	chain.AddBlock("Fifth block after Genesis")

	fmt.Println("enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("enter to continue")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
