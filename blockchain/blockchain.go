package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"goblockchain/block"
)

const (
	GENESIS  = "Genesis"
	dbPath   = "./tmp"
	LastHash = "lh"
)

type PersistentBlockchain struct {
	lastHash []byte
	Database *badger.DB
}

func InitBlockchain() *PersistentBlockchain {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	bc := &PersistentBlockchain{
		lastHash: nil,
		Database: db,
	}
	HandleError(err)
	err = db.Update(func(txn *badger.Txn) error {
		if lh, rerr := txn.Get([]byte(LastHash)); rerr == badger.ErrKeyNotFound {
			fmt.Println("PersistentBlockchain does not exist.")
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

func (bc *PersistentBlockchain) AddBlock(data string) (newBlock *block.Block) {
	prevHash := bc.lastHash
	newBlock = block.CreateBlock(data, prevHash)
	err := bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		HandleError(err)
		err = txn.Set([]byte(LastHash), newBlock.Hash)
		HandleError(err)
		bc.lastHash = newBlock.Hash
		return nil
	})
	HandleError(err)
	return
}

func GetBlockByHash(db *badger.DB, hash []byte) (b *block.Block) {
	err := db.View(func(txn *badger.Txn) error {
		get, err := txn.Get(hash)
		HandleError(err)
		valueCopy, err := get.ValueCopy([]byte{})
		deserializedBlock := block.Deserialize(valueCopy)
		b = deserializedBlock
		HandleError(err)
		return nil
	})
	HandleError(err)

	return
}

type Iterator struct {
	current *block.Block
	db      *badger.DB
}

func NewIterator(pbc *PersistentBlockchain) (it *Iterator) {
	return &Iterator{
		current: &block.Block{Hash: pbc.lastHash},
		db:      pbc.Database,
	}
}

func (it *Iterator) Next() (b *block.Block) {
	if it.current.Data == nil {
		b = GetBlockByHash(it.db, it.current.Hash)
	} else {
		b = it.current
	}

	// Genesis block's prevhash is set to nil
	if b.PrevHash == nil {
		it.current = nil
		return
	}
	it.current = GetBlockByHash(it.db, b.PrevHash)
	return
}

func (it *Iterator) HasNext() bool {
	return it.current != nil
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
