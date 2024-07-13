package main

import (
	"blockchain/src"
	"fmt"
)

func main() { // entry point function: Run command --> go run .\blockchain.go
	blockchain := src.NewBlockchain()
	blockchain.AddBlock("New Data")
	for _, block := range blockchain.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
