package chain

import (
	"fmt"
	"nexus/pkg/block"
	"nexus/pkg/db"
	"nexus/utils"
	"runtime"
)

type Chain struct {
	Blocks []*block.Block
	db     db.BadgerDB
}

func NewChain(db *db.BadgerDB) *Chain {
	chain := &Chain{
		db:     *db,
		Blocks: []*block.Block{},
	}

	chain.createGenesisBlock()

	return chain
}

func (c *Chain) createGenesisBlock() {

	value, err := c.db.Find([]byte("lh"))
	utils.HandleException(err)

	if value == nil {

		fmt.Println("creating genesis block")

		genesis := block.NewBlock(0, []byte{}, []byte{})

		serialized, err := genesis.Serialize()
		utils.HandleException(err)

		err = c.db.Save(genesis.Hash, serialized)
		utils.HandleException(err)

		err = c.db.Save([]byte("lh"), genesis.Hash)
		utils.HandleException(err)

		c.Blocks = append(c.Blocks, genesis)

	} else {
		genesisByte, err := c.db.Find(value)
		utils.HandleException(err)

		fmt.Println("genesis block found")

		genesis, err := block.Deserialize(genesisByte)
		if err != nil {
			fmt.Println(err)
			runtime.Goexit()
		}
		c.Blocks = append(c.Blocks, genesis)
	}

}

func (c *Chain) GetLatestBlock() *block.Block {
	return c.Blocks[len(c.Blocks)-1]
}

func (c *Chain) AddBlock(data []byte) error {

	var latestHash []byte
	latestHash, err := c.db.Find([]byte("lh"))
	if err != nil {
		return err
	}

	newBlock := block.NewBlock(
		c.GetLatestBlock().Index+1,
		latestHash,
		data,
	)

	serialized, err := newBlock.Serialize()
	if err != nil {
		return err
	}

	err = c.db.Save(newBlock.Hash, serialized)
	if err != nil {
		return err
	}

	err = c.db.Save([]byte("lh"), newBlock.Hash)
	if err != nil {
		return err
	}

	c.Blocks = append(c.Blocks, newBlock)

	return nil
}

// func (c *Chain) Validate() bool {
// 	for i := range c.Blocks[1:] {
// 		prevBlock := c.Blocks[i]
// 		currentBlock := c.Blocks[i+1]

// 		if !bytes.Equal(currentBlock.Hash, block.NewProofOfWork(&currentBlock).Validate()) {
// 			return false
// 		}

// 		if !bytes.Equal(currentBlock.PrevHash, prevBlock.Hash) {
// 			return false
// 		}

// 		if prevBlock.Index+1 != currentBlock.Index {
// 			return false
// 		}
// 	}

// 	return true
// }
