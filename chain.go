package main

import "fmt"

type Chain struct {
	chain      []Block
	difficulty int
}

func (c *Chain) bigBang() Block {
	return NewBlock("ancestorBlock", "")
}

func (c *Chain) GetLatestBlock() Block {
	return c.chain[len(c.chain)-1]
}

func (c *Chain) AddBlockToChain(b Block) {
	b.previousHash = c.GetLatestBlock().hash
	fmt.Printf("mine success: %s\n", b.mine(c.difficulty))
	c.chain = append(c.chain, b)
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
			fmt.Printf("chain forged: current hash %s, recomputed hash %s", b.hash, b.computeHash())
			return false
		}
	}
	return true
}

func NewChain() Chain {
	c := Chain{
		chain:      []Block{},
		difficulty: 4,
	}
	c.chain = append(c.chain, c.bigBang())
	return c
}
