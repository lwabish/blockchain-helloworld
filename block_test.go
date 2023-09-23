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

func TestMultiMine(t *testing.T) {
	b := NewBlock("lwabish", "1234")
	hash1 := b.mine(4)
	hash2 := b.mine(4)
	if hash1 != hash2 {
		t.Errorf("hash1 %s not equal to %s", hash1, hash2)
	}
}

func TestMutatingBlock(t *testing.T) {
	b := NewBlock("lwabish", "1234")
	b.mine(4)
	if b.hash != b.computeHash() {
		t.Errorf("block hash error")
	}

	b.data = "123"
	b.mine(4)
	if b.hash != b.computeHash() {
		t.Errorf("block hash error")
	}
}
