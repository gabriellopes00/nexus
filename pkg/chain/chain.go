package chain

import (
	"fmt"
	"nexus/pkg/block"
	"nexus/pkg/db"
	"nexus/utils"
	"runtime"
	"time"
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

// createGenesisBlock generate a gensis-block (first block) to the chain. It doesn't
// contain any important data. Both prevHash and data properties are set as []byte{}.
func (c *Chain) createGenesisBlock() {

	latestHash, err := c.db.Find([]byte("lh"))
	utils.HandleException(err)

	if latestHash == nil {

		fmt.Println("genesis block not found")
		fmt.Println("creating genesis block...")

		genesis := block.NewBlock(0, []byte{}, []byte{})

		serialized, err := genesis.Serialize()
		utils.HandleException(err)

		tx := c.db.NewTransaction(true)

		err = c.db.Save(genesis.Hash, serialized, tx)
		if err != nil {
			tx.Rollback()
			utils.HandleException(err)
		}

		err = c.db.Save([]byte("lh"), genesis.Hash, tx)
		if err != nil {
			tx.Rollback()
			utils.HandleException(err)
		}

		tx.Commit() // commit changes in the storage
		c.Blocks = append(c.Blocks, genesis)

		fmt.Printf("genesis block created successfully at %v\n", time.Now())

	} else {
		genesisByte, err := c.db.Find(latestHash)
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

// GetLatestBlock returns a reference to the latest block appended in the chain.
func (c *Chain) GetLatestBlock() *block.Block {
	return c.Blocks[len(c.Blocks)-1]
}

// AddBlock gets the block data and creates one. After created successfully,
// the block is registered in the storage and appended in the chain.
func (c *Chain) AddBlock(data []byte) error {

	fmt.Println("adding a new block to the chain...")

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

	tx := c.db.NewTransaction(true)

	err = c.db.Save(newBlock.Hash, serialized, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = c.db.Save([]byte("lh"), newBlock.Hash, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit() // commit ops in the storage
	c.Blocks = append(c.Blocks, newBlock)

	fmt.Printf("block added successfully at %v\n", time.Now())

	return nil
}
