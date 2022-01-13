package main

import (
	"nexus/pkg/chain"
	"nexus/pkg/cli"
	"nexus/pkg/db"
	"nexus/utils"
	"os"
)

func main() {
	defer os.Exit(0)

	blockchainDB, err := db.NewBadgerDB()
	utils.HandleException(err)

	defer blockchainDB.Conn.Close()

	blockchain := chain.NewChain(blockchainDB)

	commandLine := cli.NewCli(blockchain)
	commandLine.Run()

}
