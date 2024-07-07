package main

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func main() { // entry point function: Run command --> go run .\blockchain.go
	println("Hello world!")
}
