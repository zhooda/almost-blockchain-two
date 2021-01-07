package blockchain

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger"
)

// BlockChain contains slice of Block
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// BlockChainIterator is a structure used
// for demonstration purposes to iterate
// over our established chain
type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// InitBlockChain returns BlockChain with genesis Block
// MARK: Rename this to NewBlockChain later to comply with Go style guides
func InitBlockChain() *BlockChain {
	var lastHash []byte

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	tmpPath := filepath.Join(exPath, "tmp")
	dbPath := filepath.Join(tmpPath, "blocks")

	// Create tmp dir if doesnt already exist
	// This is a hacky workaround for badgerDB
	// not being able to create nested directories
	_ = os.MkdirAll(tmpPath, os.ModePerm)

	// Specify where database files should be stored
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath
	opts.Logger = nil

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			// If no chan exists in the database, create a
			// new chain with genesis
			fmt.Println("no existing blockchain found")
			genesis := Genesis()
			fmt.Println("genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())

			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)

			// var hashCopy []byte
			err = item.Value(func(val []byte) error {
				// Closure only called if item.Value has no error
				// Copy val to lastHash
				lastHash = append([]byte{}, val...)
				return nil
			})
			return err
		}
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

// AddBlock adds new block to the chain
// is a method for BlockChain
// data: string of data to be stored in block
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = append([]byte{}, val...)
			return nil
		})

		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		return err
	})

	Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block
	var encodedBlock []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		err = item.Value(func(val []byte) error {
			encodedBlock = append([]byte{}, val...)
			return nil
		})
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}
