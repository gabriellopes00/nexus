package main

import (
	"nexus/env"
	"nexus/pkg/chain"
)

func main() {
	blockchain := chain.NewChain(env.BLOCKCHAIN_MINING_DIFFICULTY)

	blockData := map[string]string{
		"1": "2",
		"3": "4",
		"5": "6",
	}

	blockchain.AddBlock(blockData)
	println(blockchain.IsValid())

}
