package main

func (blockchain *Blockchain) AddBlock(data string) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
