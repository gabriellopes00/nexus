package chain

import (
	"bytes"
	"encoding/gob"
	"nexus/pkg/block"
)

type Chain struct {
	blocks     []*block.Block
	difficulty int
}

func NewChain(difficulty int) *Chain {
	chain := &Chain{
		difficulty: difficulty,
	}
	chain.genGenesisBlock()
	return chain
}

func (c *Chain) genGenesisBlock() {
	block := block.NewBlock(0, []byte("0"), []byte("0"))
	c.blocks = append(c.blocks, block)
}

func (c *Chain) GetLatestBlock() *block.Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *Chain) AddBlock(data interface{}) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}

	newBlock := block.NewBlock(
		c.GetLatestBlock().Index+1,
		c.GetLatestBlock().Hash,
		buf.Bytes(),
	)

	newBlock.Mine(c.difficulty)
	c.blocks = append(c.blocks, newBlock)

	return nil
}

func (c *Chain) IsValid() bool {
	for i := range c.blocks[1:] {
		prevBlock := c.blocks[i]
		currentBlock := c.blocks[i+1]

		if !bytes.Equal(currentBlock.Hash, currentBlock.GenHash()) {
			return false
		}

		if !bytes.Equal(currentBlock.PrevHash, prevBlock.Hash) {
			return false
		}

		if prevBlock.Index+1 != currentBlock.Index {
			return false
		}
	}

	return true
}
