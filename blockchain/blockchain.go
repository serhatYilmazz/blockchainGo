package blockchain

import "goblockchain/block"

const (
	GENESIS = "Genesis"
)

type BlockChain struct {
	blocks []block.Block
}

func initBlockchain() *BlockChain {
	return &BlockChain{
		blocks: []block.Block{
			*block.CreateBlock(GENESIS, nil),
		},
	}
}

func (bc *BlockChain) genesisBlock(data string) {
	newBlock := block.CreateBlock(data, []byte(GENESIS))
	bc.blocks = append(bc.blocks, *newBlock)
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, *newBlock)
}
