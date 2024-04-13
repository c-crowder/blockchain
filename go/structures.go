package main

type Block struct {
	Timestamp     int64  // The time the block was created
	PrevBlockHash []byte // The hash of the previous block
	MyBlockHash   []byte // The hash of the current block
	AllData       []byte // The data of the transactions (body info)
}

type Blockchain struct {
	Blocks []*Block // The chain of blocks
}
