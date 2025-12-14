package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	// Replace with your actual module path
	"go_chain/blockchain"
)

type CLI struct {
	chain *blockchain.BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block <BLOCK_DATA> - Add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CLI) addBlock(data string) {
	cli.chain.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	iter := cli.chain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	// Proper exit handling for BadgerDB
	defer os.Exit(0)

	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := CLI{chain}
	cli.Run()
}
