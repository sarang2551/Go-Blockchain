package src

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Index         int64
}

/* take block fields, concatenate them, and calculate a SHA-256 hash on the concatenated combination*/
func (b *Block) SetHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
	return hash[:]
}

func NewBlock(oldBlock Block, data string) *Block {
	block := &Block{time.Now().Unix(), []byte(data), oldBlock.Hash, []byte{}, oldBlock.Index + 1}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	var genesisBlock = new(Block)
	genesisBlock.Data = []byte("GenesisBlock")
	genesisBlock.Timestamp = time.Now().Unix()
	genesisBlock.Index = 0
	genesisBlock.Hash = []byte{} // empty hash as there is no previous block
	genesisBlock.PrevBlockHash = []byte{}
	genesisBlock.SetHash()
	return genesisBlock
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		fmt.Println("Invalid index")
		return false
	}
	// check that the connection between the old and new blocks are valid
	if !bytes.Equal(oldBlock.Hash, newBlock.PrevBlockHash) {
		fmt.Println("Invalid hash connection")
		return false
	}
	// hash integrity check
	if !bytes.Equal(newBlock.SetHash(), newBlock.Hash) {
		fmt.Println("Failed hash integrity")
		return false
	}
	return true
}
