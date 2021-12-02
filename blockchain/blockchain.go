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

}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
