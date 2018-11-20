package main

import (
	"fmt"
	"strconv"

	"github.com/ObsidianRock/silsila/blockchain"
)

func main() {

	chain := blockchain.InitBlockChain()

	chain.AddBlock("Second")
	chain.AddBlock("Third")

	for _, block := range chain.Blocks {
		fmt.Printf("Data: %s, Hash: %x\n", block.Data, block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validation()))
		fmt.Println()
	}

}
