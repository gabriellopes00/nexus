package main

import (
	"fmt"
	"nexus/env"
	"nexus/pkg/chain"
	"time"
)

func main() {
	blockchain := chain.NewChain(env.BLOCKCHAIN_MINING_DIFFICULTY)
	start := time.Now()
	// fmt.Println(start)

	blockchain.AddBlock([]byte("lorem"))
	blockchain.AddBlock([]byte("ipsum"))
	blockchain.AddBlock([]byte("dolor"))
	blockchain.AddBlock([]byte("sit"))
	blockchain.AddBlock([]byte("amet"))

	fmt.Println("finished blocks creation")
	fmt.Println(time.Since(start))

	// for _, b := range blockchain.Blocks {

	// 	fmt.Printf("prev hash: %x\n", b.PrevHash)
	// 	fmt.Printf("data: %s\n", b.Data)
	// 	fmt.Printf("hash: %x\n", b.Hash)

	// 	pow := block.NewProofOfWork(b)
	// 	fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println("")

	// }
	// println(blockchain.IsValidChain())

}
