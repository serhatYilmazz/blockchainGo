package block

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

type Proof struct {
	Block      *Block
	Target     big.Int
	Difficulty int
}

const (
	HashSpace = 512
)

func NewProof(b *Block, difficulty int) *Proof {
	one := big.NewInt(1)
	target := one.Lsh(one, uint(HashSpace-difficulty))
	return &Proof{
		Block:  b,
		Target: *target,
		Difficulty: difficulty,
	}
}

func (p *Proof) Generate() ([]byte, int) {
	nonce := 0
	var hash [64]byte
	for nonce < math.MaxInt {
		data := p.createData(nonce)
		hash = deriveHash(data)
		fmt.Printf("\r%x", hash)
		isValid := p.validateHash(hash)
		if isValid {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return hash[:], nonce
}

func (p *Proof) validateHash(hash [64]byte) bool {
	var intHash big.Int
	intHash.SetBytes(hash[:])

	if intHash.Cmp(&p.Target) == -1 {
		return true
	}
	return false
}

func deriveHash(data []byte) [64]byte {
	return sha512.Sum512(data)
}

func (p *Proof) createData(nonce int) (data []byte) {
	data = bytes.Join([][]byte{
		p.Block.PrevHash,
		p.Block.Data,
		Int64ToByteSlice(int64(nonce)),
		Int64ToByteSlice(int64(p.Difficulty)),
	},
		[]byte{})
	return
}

func Int64ToByteSlice(d int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, d)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}
