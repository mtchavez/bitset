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

func Test_Set(t *testing.T) {
	b := New(5)
	b.Set(4)
	if b.bits[4>>logWordSize] != 1<<4 {
		t.Errorf("Expected bit 4 to have been set but got %b", b.bits)
	}
}

func Test_Set_greaterThanLength(t *testing.T) {
	b := New(4)
	b.Set(6)
	if b.bits[6>>logWordSize] != 1<<6 {
		t.Errorf("Expected bit 5 to have been set but got %b", b.bits)
	}

	if b.length != 7 {
		t.Errorf("Expected length to be increased to 7 but got %+v", b.length)
	}
}

func Test_Set_growBits(t *testing.T) {
	b := New(4)
	b.Set(32)
	b.Set(64)
	if b.bits[32>>wordSize] != 1<<32 {
		t.Errorf("Expected bit 32 to have been set but got %b", b.bits)
	}

	if b.bits[64>>logWordSize] != 1<<(64&(wordSize-1)) {
		t.Errorf("Expected bit 64 to have been set but got %b", b.bits[64>>wordSize])
	}

	if b.length != 65 {
		t.Errorf("Expected length to be increased to 65 but got %+v", b.length)
	}
}

func Test_growBits(t *testing.T) {
	b := New(4)
	grown := b.growBits(64)
	if !grown {
		t.Errorf("Expected bits to be grown")
	}
	if b.length != 65 {
		t.Errorf("Expected length to be increased to 65 but got %+v", b.length)
	}
}

func Test_Test_largerThanLength(t *testing.T) {
	b := New(4)
	found := b.Test(6)
	if found {
		t.Errorf("Expected to not find a larger value then length")
	}
}

func Test_Test_notFound(t *testing.T) {
	b := New(4)
	found := b.Test(3)
	if found {
		t.Errorf("Expected to not find unset 3")
	}
}

func Test_Test_found(t *testing.T) {
	b := New(4)
	b.Set(3)
	found := b.Test(3)
	if !found {
		t.Errorf("Expected to find 3")
	}
}
