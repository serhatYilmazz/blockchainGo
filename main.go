package main

import (
	"fmt"
	"goblockchain/blockchain"
)

func main() {
	initBlockchain := blockchain.InitBlockchain()
	//initBlockchain.AddBlock("This is is")
	//initBlockchain.AddBlock("This is Serhatt")
	//initBlockchain.AddBlock("This is Yilmaz")
	iterator := blockchain.NewIterator(initBlockchain)
	for iterator.HasNext() {
		fmt.Printf("%+v\n", iterator.Next().Print())
	}
}
