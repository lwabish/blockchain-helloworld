package main

import "fmt"

type Chain struct {
	chain            []Block
	transactionPools []Transaction
	difficulty       int
	minerReward      int
}

func (c *Chain) bigBang() Block {
	return NewBlock(c.transactionPools, "")
}

func (c *Chain) GetLatestBlock() Block {
	return c.chain[len(c.chain)-1]
}

// func (c *Chain) AddBlockToChain(b Block) {
// 	b.previousHash = c.GetLatestBlock().hash
// 	fmt.Printf("mine success: %s\n", b.mine(c.difficulty))
// 	c.chain = append(c.chain, b)
// }

func (c *Chain) MineTransactionPool(minerAddr string) {
	minerRewardTransaction := NewTransaction("", minerAddr, c.minerReward)
	c.transactionPools = append(c.transactionPools, minerRewardTransaction)

	b := NewBlock(c.transactionPools, c.GetLatestBlock().hash)
	b.mine(c.difficulty)

	c.chain = append(c.chain, b)
	c.transactionPools = []Transaction{}
}

func (c *Chain) Validate() bool {
	if len(c.chain) == 1 {
		ancestorBlock := c.chain[0]
		if ancestorBlock.computeHash() == ancestorBlock.hash {
			return true
		} else {
			return false
		}
	}
	for k, b := range c.chain[1:] {
		lastBlock := c.chain[k]
		if b.previousHash != lastBlock.hash {
			fmt.Println("chain broken")
			return false
		}
		if b.hash != b.computeHash() {
			fmt.Printf("chain forged: current hash %s, recomputed hash %s\n", b.hash, b.computeHash())
			return false
		}
	}
	return true
}

func NewChain() Chain {
	c := Chain{
		difficulty:  4,
		minerReward: 50,
	}
	c.chain = append(c.chain, c.bigBang())
	return c
}
