package block

import (
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
	printed := block.Print()

	// If they have three fields and specifially they are: ...
	if !compileFields.MatchString(printed) {
		t.Errorf("Error")
	}

}

func TestBlock_DeriveHash(t *testing.T) {

}
