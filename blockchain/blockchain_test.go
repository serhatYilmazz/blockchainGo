package blockchain

import (
	"testing"
)

func TestBlockChain_genesisBlock(t *testing.T) {
	bc := initBlockchain()
	if len(bc.blocks) != 1 {
		t.Errorf("There should be exactly one block in blockchain, but there is %d", len(bc.blocks))
	}
}

func TestBlockChain_AddBlock(t *testing.T) {
	s := "This is an example block data"
	bc := initBlockchain()
	bc.AddBlock(s)
	if len(bc.blocks) != 2 {
		t.Errorf("There should be exactly one block in blockchain, but there is %d", len(bc.blocks))
	}
}
