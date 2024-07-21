package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	//"github.com/joho/godotenv"
	"github.com/sarang2551/Go-Blockchain/src"
)

func main() { // entry point function: Run command --> go run .\main.go

	done := make(chan struct{}) // creating a channel to prevent race condition
	go func() {

		bc := src.GetBlockchainInstance()
		spew.Dump(bc.Blocks[0])
		fmt.Println("Length of blockchain: ", len(bc.Blocks))
	}()
	close(done)
	log.Fatal(src.Run())

}
