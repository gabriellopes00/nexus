package chain

import (
	"nexus/pkg/block"
)

type Chain struct {
	Blocks     []*block.Block
	difficulty int
}

func NewChain(difficulty int) *Chain {
	chain := &Chain{
		difficulty: difficulty,
	}
	chain.createGenesisBlock()
	return chain
}

func (c *Chain) createGenesisBlock() {
	block := block.NewBlock(0, []byte{}, []byte{})
	c.Blocks = append(c.Blocks, block)
}

func (c *Chain) GetLatestBlock() *block.Block {
	return c.Blocks[len(c.Blocks)-1]
}

func (c *Chain) AddBlock(data []byte) error {
	// var buf bytes.Buffer
	// encoder := gob.NewEncoder(&buf)
	// err := encoder.Encode(data)
	// if err != nil {
	// 	return err
	// }

	newBlock := block.NewBlock(
		// c.GetLatestBlock().Index+1,
		1,
		c.GetLatestBlock().Hash,
		data,
	)

	// newBlock.Mine(c.difficulty)
	c.Blocks = append(c.Blocks, newBlock)

	return nil
}

// func (c *Chain) IsValidChain() bool {
// 	for i := range c.Blocks[1:] {
// 		prevBlock := c.Blocks[i]
// 		currentBlock := c.Blocks[i+1]

// 		if !bytes.Equal(currentBlock.Hash, currentBlock.DriveHash()) {
// 			return false
// 		}

// 		if !bytes.Equal(currentBlock.PrevHash, prevBlock.Hash) {
// 			return false
// 		}

// 		// if prevBlock.Index+1 != currentBlock.Index {
// 		// 	return false
// 		// }
// 	}

// 	return true
// }
