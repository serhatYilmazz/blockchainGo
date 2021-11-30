package blockchain

import "goblockchain/block"

const (
	GENESIS = "Genesis"
)

type Blockchain struct {
	blocks []block.Block
}

func initBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []block.Block{
			*block.CreateBlock(GENESIS, nil),
		},
	}
}

func (bc *Blockchain) genesisBlock(data string) {
	newBlock := block.CreateBlock(data, []byte(GENESIS))
	bc.blocks = append(bc.blocks, *newBlock)
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, *newBlock)
}
