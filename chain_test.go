package main

import "testing"

func TestNewChain(t *testing.T) {
	c := NewChain()
	t.Log(c)
}

func TestMineWithChain(t *testing.T) {
	c := NewChain()
	c.AddTransaction(NewTransaction("a1", "a2", 100))
	c.AddTransaction(NewTransaction("a2", "a3", 25))
	c.MineTransactionPool("a4")
	t.Log(c)
}

func TestValidateChain(t *testing.T) {
	c := NewChain()
	t.Logf("validate chain with only ancestor block result: %v", c.Validate())
	c.MineTransactionPool("a1")
	t.Log(c)
	t.Logf("validate chain after add some blocks result: %v", c.Validate())

	c.chain[1].transactions = []Transaction{NewTransaction("a2", "a1", 45)}
	t.Logf("validate chain after forge the block data result: %v", c.Validate())
	c.chain[1].mine(4)
	t.Logf("validate chain after forge the block hash result: %v", c.Validate())

}
