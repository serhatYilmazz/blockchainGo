package blockchain

import (
	"fmt"
	"goblockchain/block"
)

const (
	GENESIS = "Genesis"
)

type Blockchain struct {
	Blocks []block.Block
}

func initBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []block.Block{
			*block.CreateBlock(GENESIS, nil),
		},
	}
}

func (bc *Blockchain) AddBlock(data string) error {
	if len(bc.Blocks) != 1 {
		return fmt.Errorf("Initialize blockchain before add a new block")
	}
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, *newBlock)
	return nil
}
