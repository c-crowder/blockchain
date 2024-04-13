package main

import (
	"bytes"         // need to convert data into byte in order to be sent on the network
	"crypto/sha256" // crypto library to hash the data
	"strconv"       // convert int to string
	"time"          // get the time for timestamp
)

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PrevBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
