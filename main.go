package main

import (
	"fmt"

	"github.com/gabriel1680/blockchain/core"
)

func main() {
	chain := core.InitBockChain()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	for _, block := range chain.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
	}
}
