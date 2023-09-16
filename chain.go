package main

type Chain struct {
	chain []Block
}

func (c *Chain) bigBang() Block {
	return NewBlock("ancestorBlock", "")
}

func NewChain() Chain {
	c := Chain{
		chain: []Block{},
	}
	c.chain = append(c.chain, c.bigBang())
	return c
}
