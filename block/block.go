package block

import (
	"bytes"
	"crypto/sha512"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	blockInfo := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha512.Sum512(blockInfo)
	b.Hash= hash[:]
}

func (b *Block) Print() string {
	return fmt.Sprintf("Data: %s\nHash:%s\nprevHash: %s\n\n", string(b.Data), string(b.Hash), string(b.PrevHash))
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     nil,
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}
