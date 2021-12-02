package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"testing"
)

func TestBlockChain_genesisBlock(t *testing.T) {
	bc := InitBlockchain()
	fmt.Println(len(bc.lastHash))
	if len(bc.lastHash) != 64 {
		t.Errorf("There should be exactly one block in blockchain")
	}
}

func TestBlockChain_AddBlock(t *testing.T) {
	s := "This is an example block data"
	bc := InitBlockchain()
	block := bc.AddBlock(s)
	blockDb := bc.GetBlockByHash(bc.lastHash)
	if block.Nonce != blockDb.Nonce {
		t.Errorf("AddBlock does not run properly")
	}
	deleteFromBlockchain(bc, block.Hash)
}

func deleteFromBlockchain(bc *Blockchain, hashToDelete []byte) {
	err := bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Delete(hashToDelete)
		HandleError(err)
		return nil
	})
	HandleError(err)
}