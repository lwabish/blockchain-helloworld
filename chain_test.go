package main

import "testing"

func TestNewChain(t *testing.T) {
	c := NewChain()
	t.Log(c)
}

func TestAddBlock(t *testing.T) {
	c := NewChain()
	b := NewBlock("transfer10Yuan", "")
	b1 := NewBlock("transferAnother10Yuan", "")
	c.AddBlockToChain(b)
	c.AddBlockToChain(b1)
	t.Log(c)
}

func TestValidateChain(t *testing.T) {
	c := NewChain()
	t.Logf("validate chain with only ancestor block result: %v", c.Validate())

	b := NewBlock("transfer10Yuan", "")
	b1 := NewBlock("transferAnother10Yuan", "")
	c.AddBlockToChain(b)
	c.AddBlockToChain(b1)
	t.Log(c)
	t.Logf("validate chain after add some blocks result: %v", c.Validate())

	c.chain[1].data = "transfer100Yuan"
	t.Logf("validate chain after forge the block result: %v", c.Validate())
}
