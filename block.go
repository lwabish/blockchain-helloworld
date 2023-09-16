package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	data         string
	previousHash string
	hash         string
}

func (b *Block) computeHash() string {
	sum := sha256.Sum256([]byte(b.data + b.previousHash))
	return fmt.Sprintf("%x", sum)
}

func NewBlock(data string, hash string) Block {
	b := Block{
		data:         data,
		previousHash: hash,
	}
	b.hash = b.computeHash()
	return b
}
