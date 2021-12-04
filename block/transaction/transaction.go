package transaction

import (
	"bytes"
	"encoding/gob"
	"goblockchain/util"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxInput struct {
	ID        []byte
	OutRef    int
	ScriptSig [][]byte // 0. index: signature
	// 1. index: public key
}

type TxOutput struct {
	Value        int
	ScriptPubKey []byte // public key hash
}

func CoinbaseTransaction(sig, pubKey string) *Transaction {
	pubKeyHash := util.DeriveHash([]byte(pubKey))
	txIn := TxInput{
		ID:        []byte{},
		OutRef:    -1,
		ScriptSig: [][]byte{[]byte(sig), []byte(pubKey)},
	}
	txOut := TxOutput{
		Value:        100,
		ScriptPubKey: pubKeyHash[:],
	}

	tx := &Transaction{
		ID:      nil,
		Inputs:  []TxInput{txIn},
		Outputs: []TxOutput{txOut},
	}

	tx.SetID()
	return tx
}

func (tx *Transaction) IsCoinBaseTransaction() bool {
	return tx.Inputs[0].OutRef == -1 && len(tx.Inputs[0].ID) == 0
}

func (tx *Transaction) SetID() {
	serializedTx := tx.Serialize()
	hash := util.DeriveHash(serializedTx)
	hashSlice := hash[:]
	tx.ID = hashSlice
}

func (tx *Transaction) Serialize() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)
	util.HandleError(err)

	return buf.Bytes()
}
