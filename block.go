package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	AnswerPrefix = "0"
)

type Block struct {
	data         string
	previousHash string
	hash         string
	nonce        int
}

func (b *Block) generateAnswer(diffculty int) string {
	return strings.Repeat(AnswerPrefix, diffculty)
}

func (b *Block) mine(difficulty int) string {
	for {
		if b.hash[:difficulty] != b.generateAnswer(difficulty) {
			b.nonce++
			b.hash = b.computeHash()
		} else {
			return b.hash
		}
	}
}

func (b *Block) computeHash() string {
	sum := sha256.Sum256([]byte(b.data + b.previousHash + fmt.Sprintf("%d", b.nonce)))
	return fmt.Sprintf("%x", sum)
}

// fixme: hash seems useless
func NewBlock(data string, hash string) Block {
	b := Block{
		data:         data,
		previousHash: hash,
	}
	b.hash = b.computeHash()
	return b
}
