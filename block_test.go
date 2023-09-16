package main

import "testing"

func TestNewBlock(t *testing.T) {
	b := NewBlock("WUBOWEN", "1234")
	t.Log(b)
}
