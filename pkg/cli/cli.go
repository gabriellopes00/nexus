package cli

import (
	"flag"
	"fmt"
	"nexus/pkg/block"
	"nexus/pkg/chain"
	"os"
	"runtime"
	"strconv"
)

type CLI struct {
	blockchain *chain.Chain
}

func NewCli(blockchain *chain.Chain) *CLI {
	return &CLI{
		blockchain: blockchain,
	}
}

func (CLI) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("add -block BLOCK_DATA - add a new block to the chain")
	fmt.Println("print - prints the entire blockchain")
}

func (cli *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		runtime.Goexit() // exit the go routine - importante once badger db has its own garbage collector, and needs to be
		// stopped properly
	}
}

func (cli *CLI) AddBlock(data string) {
	err := cli.blockchain.AddBlock([]byte(data))
	if err != nil {
		fmt.Println(err)
		runtime.Goexit()
	}

	fmt.Println("Block Added Successfully !")
}

func (cli *CLI) PrintChain() {
	iter := cli.blockchain.Iterator()

	for {

		b, err := iter.Next()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while printing the chain")
			runtime.Goexit()
		}

		fmt.Printf("Previous block hash: %x\n", b.PrevHash)
		fmt.Printf("Block data: %s\n", string(b.Data))
		fmt.Printf("Block hash: %x\n", b.Hash)

		fmt.Println()

		pow := block.NewProofOfWork(b)
		fmt.Printf("Is valid pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(b.PrevHash) == 0 {
			break
		}

	}
}

func (cli *CLI) Run() {
	cli.PrintUsage()
	fmt.Println("")

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		{
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil {
				fmt.Println(err)
				runtime.Goexit()
			}
		}
	case "print":
		{
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil {
				fmt.Println(err)
				runtime.Goexit()
			}
		}
	default:
		{
			cli.PrintUsage()
			runtime.Goexit()
		}
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}

		cli.AddBlock(*addBlockData)
		return
	}

	if printChainCmd.Parsed() {
		cli.PrintChain()
		return
	}
}
