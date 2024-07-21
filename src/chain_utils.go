package src

import (
	"fmt"
	"sync"
)

type Blockchain struct {
	Blocks []*Block
}

type BlockchainManager struct {
	instance *Blockchain
	once     sync.Once
}

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
func (bcm *BlockchainManager) GetInstance() *Blockchain {
	bcm.once.Do(func() {
		bcm.instance = &Blockchain{}
	})
	return bcm.instance
}

func NewBCM() *BlockchainManager {
	return &BlockchainManager{}
}
