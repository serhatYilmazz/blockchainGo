package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"goblockchain/block"
)

const (
	GENESIS = "Genesis"
	dbPath  = "./tmp/blockchain"
	LastHash = "lh"
)

type Blockchain struct {
	lastHash []byte
	Database *badger.DB
}

type Iterator struct {
	currentHash []byte
	db       *badger.DB
}

func InitBlockchain() *Blockchain {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	bc := &Blockchain{
		lastHash: nil,
		Database: db,
	}
	HandleError(err)
	err = db.Update(func(txn *badger.Txn) error {
		if lh, rerr := txn.Get([]byte(LastHash)); rerr == badger.ErrKeyNotFound {
			fmt.Println("Blockchain does not exist.")
			genesis := genesis()
			serializedBlock := genesis.Serialize()
			bc.lastHash = genesis.Hash
			rerr := txn.Set(genesis.Hash, serializedBlock)
			HandleError(rerr)

			rerr = txn.Set([]byte(LastHash), genesis.Hash)
			HandleError(rerr)
			return nil
		} else {
			fmt.Println("Last hash is ", lh)
			bc.lastHash, rerr = lh.ValueCopy([]byte{})
			HandleError(rerr)
			return nil
		}
	})
	HandleError(err)

	return bc
}

func genesis() *block.Block {
	createdBlock := block.CreateBlock(GENESIS, nil)
	return createdBlock
}

func (bc *Blockchain) AddBlock(data string) *block.Block {
	prevHash := bc.lastHash
	newBlock := block.CreateBlock(data, prevHash)
	err := bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		HandleError(err)
		bc.lastHash = newBlock.Hash
		return nil
	})
	HandleError(err)
	return newBlock
}

func (bc *Blockchain) GetBlockByHash(hash []byte) *block.Block {
	var b *block.Block
	err := bc.Database.View(func(txn *badger.Txn) error {
		get, err := txn.Get(hash)
		HandleError(err)
		valueCopy, err := get.ValueCopy([]byte{})
		deserializedBlock := block.Deserialize(valueCopy)
		b = deserializedBlock
		HandleError(err)
		return nil
	})
	HandleError(err)

	return b
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
