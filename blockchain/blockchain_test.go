package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"testing"
)

func TestBlockChain_genesisBlock(t *testing.T) {
	bc := InitBlockchain()
	defer bc.Database.Close()
	fmt.Println(len(bc.lastHash))
	if len(bc.lastHash) != 64 {
		t.Errorf("There should be exactly one block in blockchain")
	}
	cleanDb(bc)
}

func TestBlockChain_AddBlock(t *testing.T) {
	s := "This is an example block data"
	bc := InitBlockchain()
	defer bc.Database.Close()
	block := bc.AddBlock(s)
	blockDb := GetBlockByHash(bc.Database, bc.lastHash)
	if block.Nonce != blockDb.Nonce {
		t.Errorf("AddBlock does not run properly")
	}
	cleanDb(bc)
}

func TestIterator_Iterator_Next_HasNext(t *testing.T) {
	d1 := "This is an example block data d1"
	d2 := "This is an example block data d2"
	d3 := "This is an example block data d3"
	bc := InitBlockchain()
	defer bc.Database.Close()
	bc.AddBlock(d1)
	bc.AddBlock(d2)
	bc.AddBlock(d3)

	it := NewIterator(bc)
	counter := 0
	for it.HasNext() {
		counter++
		it.Next().Print()
	}
	if counter != 4 {
		t.Errorf("Total number of blocks should be %d, but it is %d", 4, counter)
	}
	cleanDb(bc)
}

		err := txn.Delete(hashToDelete)
		HandleError(err)
		return nil
	})
	HandleError(err)
}
