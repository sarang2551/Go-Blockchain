package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/sarang2551/Go-Blockchain/src"
)

func main() { // entry point function: Run command --> go run .\main.go
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		genesisBlock := src.NewGenesisBlock()
		spew.Dump(genesisBlock)
		bc := src.NewBCM().GetInstance()
		bc.AddBlock(genesisBlock)
		fmt.Println("Length of blockchain: ", len(bc.Blocks))
	}()
	log.Fatal(src.Run())

}
