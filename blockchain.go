package main

import (
	"blockchain/src"
)

func main() { // entry point function: Run command --> go run .\blockchain.go
	data := "Some data for the block"
	prevBlockHash := []byte{0x01, 0x02, 0x03, 0x04} // Example previous block hash
	newBlock := src.NewBlock(data, prevBlockHash)
	print(newBlock)
}
