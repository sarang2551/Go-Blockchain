package src

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Blockchain struct {
	Blocks []*Block
}

var blockchainInstance *Blockchain

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(newBlock *Block) {
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) GetLastBlock() (*Block, error) {
	if len(bc.Blocks) == 0 {
		return nil, fmt.Errorf("Blockchain is empty")
	}
	return bc.Blocks[len(bc.Blocks)-1], nil
}

// Getting the Singleton instance of the blockchain
func GetBlockchainInstance() *Blockchain {
	if blockchainInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if blockchainInstance == nil {
			fmt.Println("Creating single instance now.")
			blockchainInstance = &Blockchain{}
			fmt.Println("Adding genesis block...")
			blockchainInstance.AddBlock(NewGenesisBlock())
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return blockchainInstance
}
