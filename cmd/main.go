package main

import (
	"fmt"
	"nexus/pkg/chain"
	"nexus/pkg/cli"
	"nexus/pkg/db"
	"os"
	"runtime"
)

func main() {
	defer os.Exit(0)

	blockchainDB, err := db.NewBadgerDB()
	if err != nil {
		fmt.Println(err)
		runtime.Goexit()
	}

	defer blockchainDB.Conn.Close()

	blockchain := chain.NewChain(blockchainDB)

	commandLine := cli.NewCli(blockchain)
	commandLine.Run()

}
