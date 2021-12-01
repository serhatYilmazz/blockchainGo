package block

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewProof(t *testing.T) {
	block := CreateBlock("dummy data", []byte("dummy prev Hash"))
	proof := NewProof(block)
	if proof.Block.Print() != block.Print() {
		t.Errorf("Newly created proof is ntot true")
	}
}

func TestProof_Generate(t *testing.T) {
	block := CreateBlock("dummy data", []byte("dummy prev Hash"))
	proof := NewProof(block)
	hash, generatedNonce := proof.Generate()

	if proof.Block.Nonce != generatedNonce {
		t.Errorf("Generated and Asssigned nonces are not equal.")
	}

	if bytes.Compare(proof.Block.Hash, hash) != 0 {
		t.Errorf("Generated and Asssigned nonces are not equal.")
	}

	fmt.Printf("%s\n", block.Print())
}

func TestInt64ToByteSlice(t *testing.T) {
	slice := Int64ToByteSlice(int64(4))
	if slice[7] != 4 {
		t.Errorf("Byte slice conversion does not work right")
	}
}
