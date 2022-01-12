package chain

import (
	"fmt"
	"nexus/pkg/block"
	"nexus/pkg/db"
)

type ChainIterator struct {
	CurrentHash []byte
	db          db.BadgerDB
}

func (c *Chain) Iterator() *ChainIterator {
	return &ChainIterator{
		CurrentHash: c.GetLatestBlock().Hash,
		db:          c.db,
	}
}

func (i *ChainIterator) Next() (*block.Block, error) {
	var b *block.Block

	data, err := i.db.Find(i.CurrentHash)
	if err != nil {
		if data == nil {
			return nil, nil
		}

		return nil, err
	}

	b, err = block.Deserialize(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	i.CurrentHash = b.PrevHash

	return b, nil
}
