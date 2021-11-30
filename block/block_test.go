package block

import (
	"crypto/sha512"
	"regexp"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	prevHash := []byte("dummy hash")
	block := CreateBlock("dummy data", prevHash)
	if block.PrevHash == nil {
		t.Errorf("Previous hash value should not be nil")
	}

	if block.Hash == nil {
		t.Errorf("Block's hash value should not be nil")
	}

	if block.Data == nil {
		t.Errorf("Block's data value should not be nil")
	}
}

func TestBlock_Print(t *testing.T) {
	prevHash := []byte("dummy hash")
	block := CreateBlock("dummy data", prevHash)
	compileFields, _ := regexp.Compile("(Data:).+\n(Hash:).+\n(prevHash:).+")
	compileNewLines, _ := regexp.Compile("[\n]")
	printed := block.Print()

	// If they have three fields and specifially they are: ...
	if len(compileNewLines.FindAllString(printed, -1)) != 3 && !compileFields.MatchString(printed) {
		t.Errorf("Error")
	}

}

func TestBlock_DeriveHash(t *testing.T) {
	bytes := sha512.Sum512([]byte("random"))
	b := &Block{
		Data:     []byte{},
		PrevHash: bytes[:],
	}
	b.DeriveHash()
	if b.Hash == nil {
		t.Errorf("Block's hash should not be nil")
	}

}
