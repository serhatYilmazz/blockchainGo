package blockchain

import "goblockchain/block"

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

func (bc *Blockchain) genesisBlock(data string) {
	newBlock := block.CreateBlock(data, []byte(GENESIS))
	bc.Blocks = append(bc.Blocks, *newBlock)
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, *newBlock)
}
