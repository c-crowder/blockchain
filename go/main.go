package main

import (
	"fmt"
)

func main() {
	newblockchain := NewBlockchain() // Initialize the blockchain
	// create two blocks and add two transactions
	newblockchain.AddBlock("First transaction")
	newblockchain.AddBlock("Second transaction")
	// Print all the blocks and their content
	for i, block := range newblockchain.Blocks {
		fmt.Printf("Block ID: %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block: %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous block: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.AllData)
		fmt.Println()
	}
}
