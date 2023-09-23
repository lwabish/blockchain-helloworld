package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	AnswerPrefix = "0"
)

type Block struct {
	transactions []Transaction
	previousHash string
	hash         string
	nonce        int
	timestamp    int64
}

func (b *Block) generateAnswer(diffculty int) string {
	return strings.Repeat(AnswerPrefix, diffculty)
}

func (b *Block) mine(difficulty int) string {
	for {
		// 每次都重算一下，确保幂等性。
		// 如果不重算，数据被修改后，可能hash不一致
		b.hash = b.computeHash()
		if b.hash[:difficulty] != b.generateAnswer(difficulty) {
			b.nonce++
			b.hash = b.computeHash()
		} else {
			return b.hash
		}
	}
}

func (b *Block) computeHash() string {
	transactionStrings, err := json.Marshal(b.transactions)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256([]byte(string(transactionStrings) + b.previousHash + fmt.Sprintf("%d%d", b.nonce, b.timestamp)))
	return fmt.Sprintf("%x", sum)
}

func NewBlock(transactions []Transaction, hash string) Block {
	b := Block{
		transactions: transactions,
		previousHash: hash,
		timestamp:    time.Now().Unix(),
	}
	b.hash = b.computeHash()
	return b
}
