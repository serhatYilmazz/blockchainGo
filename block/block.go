package block

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

const (
	Difficulty = 12
)

func (b *Block) Print() string {
	return fmt.Sprintf("Data: %s\nHash:%x\nPrevHash: %x\nNonce: %d\n", string(b.Data), b.Hash, b.PrevHash, b.Nonce)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     nil,
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}
	proof := NewProof(block, Difficulty)
	block.Hash, block.Nonce = proof.Generate()

	return block
}

func (b *Block) Serialize() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println(err)
	}

	return buf.Bytes()
}

func Deserialize(data []byte) (b *Block) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Println(err)
	}
	return
}