package main

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	b := NewBlock("lwabish", "1234")
	t.Log(b)
}

func TestGenerateAnswer(t *testing.T) {
	b := NewBlock("lwabish", "1234")
	t.Log(b.generateAnswer(3))
}

func TestStringCut(t *testing.T) {
	t.Log("abcde"[:3])
}

func TestMine(t *testing.T) {
	b := NewBlock("lwabish", "1234")
	t.Logf("\nmine success: %s", b.mine(3))
}
