package main

import "testing"

func TestNewChain(t *testing.T) {
	c := NewChain()
	t.Log(c)
}

func TestAddBlock(t *testing.T) {
	c := NewChain()
	b := NewBlock("transfer10yuan", "")
	c.AddBlockToChain(b)
	t.Log(c)
}
