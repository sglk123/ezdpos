package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Block simple struct
type Block struct {
	Height    int
	Timestamp string
	PreHash   string
	Hash      string
	RawData   string
	Validator string
}

// Delegate can be each node
type Delegate struct {
	Id    string
	Votes int
}

// GenerateGenesisBlock is used to propose first block
func GenerateGenesisBlock() *Block {
	block := Block{
		Height:    0,
		Timestamp: time.Now().String(),
		PreHash:   "",
		Hash:      "",
		RawData:   "",
		Validator: "root",
	}
	block.Hash = CalculateBlockHash(&block)
	return &block
}

func GenerateNextBlock(lastBlock *Block, validator string) *Block {
	block := Block{
		Height:    lastBlock.Height + 1,
		Timestamp: time.Now().String(),
		PreHash:   lastBlock.Hash,
		Hash:      "",
		RawData:   "",
		Validator: validator,
	}
	block.Hash = CalculateBlockHash(&block)
	return &block
}
func CalculateBlockHash(block *Block) string {
	hashData := strconv.Itoa(block.Height) + block.Timestamp + block.RawData + block.PreHash
	h := sha256.New()
	h.Write([]byte(hashData))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func main() {
	var blockChain []*Block
	// init 6 members and rand votes
	validatorList := make([]Delegate, 6)
	for i := 0; i < 6; i++ {
		validatorList[i].Id = strconv.Itoa(i)
		rand.Seed(time.Now().UnixNano())
		validatorList[i].Votes = rand.Intn(100)
		fmt.Printf("validator id is %v, vote is %v \n", i, validatorList[i].Votes)
		time.Sleep(1 * time.Second)
	}
	curBlock := GenerateGenesisBlock()
	blockChain = append(blockChain, curBlock)
	// select two winners to generate block
	validators := GetValidators(validatorList)

	// generate block by winners
	for _, value := range validators {
		curBlock = GenerateNextBlock(curBlock, value.Id)
		blockChain = append(blockChain, curBlock)
	}

	for i := 0; i < len(blockChain); i++ {
		fmt.Printf("block height is %v, block validator is %v \n", blockChain[i].Height, blockChain[i].Validator)
	}
}

// GetValidators is to get top 2 voters
func GetValidators(list []Delegate) []Delegate {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list)-1; j++ {
			if list[i].Votes < list[j].Votes {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list[:2]
}
