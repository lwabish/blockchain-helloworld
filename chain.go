package main

type Chain struct {
	chain []Block
}

func (c *Chain) bigBang() Block {
	return NewBlock("ancestorBlock", "")
}

func (c *Chain) GetLatestBlock() Block {
	return c.chain[len(c.chain)-1]
}

func (c *Chain) AddBlockToChain(b Block) {
	b.previousHash = c.GetLatestBlock().hash
	b.hash = b.computeHash()
	c.chain = append(c.chain, b)
}

func NewChain() Chain {
	c := Chain{
		chain: []Block{},
	}
	c.chain = append(c.chain, c.bigBang())
	return c
}
