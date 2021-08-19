package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// Convert int to byte array
func Int2Byte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}

// Block structure
type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	TimeStamp     int64
	TargetBits    int64
	Nonce         int64
	MerkeRoot     []byte
	Data          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		TimeStamp:     time.Now().Unix(),
		TargetBits:    10,
		Nonce:         5,
		MerkeRoot:     []byte{},
		Data:          []byte(data)}
	block.SetHash()

	return block

}

// int -> byte

func (block *Block) SetHash() {
	tmp := [][]byte{
		// Function to convert int into byte array
		Int2Byte(block.Version),
		block.PrevBlockHash,
		Int2Byte(block.TimeStamp),
		block.MerkeRoot,
		Int2Byte(block.Nonce),
		block.Data}
	// Connect block fields into a slice and use [] byte {} to connect

	data := bytes.Join(tmp, []byte{})

	// Calculate the value of hash
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}

type BlockChain struct {
	// Use slices to save blocks for blockchain simulation
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	// Create a blockchain
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
	// Block crossing prevention
	if len(bc.blocks) <= 0 {
		os.Exit(1)
	}
	lastBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Test the first BTC")
	bc.AddBlock("Test the second EOS")
	bc.AddBlock("Test the third")
	bc.AddBlock("Test the fourth")

	for i, block := range bc.blocks {
		fmt.Println("=========block num:", i)
		fmt.Println("data", string(block.Data))
		fmt.Println("Version:", block.Version)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("TimeStamp:%d\n", block.TimeStamp)
		fmt.Printf("MerkeRoot:%x\n", block.MerkeRoot)
		fmt.Printf("None:%d\n", block.Nonce)
	}
}
