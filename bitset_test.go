package bitset

import (
	"math"
	"testing"
)

func Test_wordsNeeded_lessThanMax(t *testing.T) {
	needed := wordsNeeded(2)
	if needed != 1 {
		t.Errorf("Expected 1 but got %+v", needed)
	}
	needed = wordsNeeded(80)
	if needed != 2 {
		t.Errorf("Expected 2 but got %+v", needed)
	}
}

func Test_wordsNeeded_greaterThanMax(t *testing.T) {
	needed := wordsNeeded(math.MaxUint64 - 3)
	if needed != 288230376151711743 {
		t.Errorf("Expected 288230376151711743 but got %+v", needed)
	}
}

func Test_New(t *testing.T) {
	b := New(500)
	if b == nil {
		t.Errorf("Should have created a bitset but got nil")
	}
	if b.length != 500 {
		t.Errorf("Should have gotten a length of 500 but got %+v", b.length)
	}
	if uint64(len(b.bits)) != wordsNeeded(500) {
		t.Errorf("Length of bits should have been %+v but got %+v", wordsNeeded(500), b.bits)
	}
}
