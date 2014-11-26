package bitset

import "math"

// Bitset struct to wrap underlying bits
type Bitset struct {
	length uint64
	bits   []uint64
}

const (
	wordSize    uint64 = 64
	logWordSize uint64 = 6
)

func wordsNeeded(i uint64) uint64 {
	if i > math.MaxUint64-wordSize+1 {
		return math.MaxUint64 >> logWordSize
	}
	return (i + (wordSize - 1)) >> logWordSize
}

// New will make a bitset with the length set
// and the necessary bits allocated for the
// requested length
func New(length uint64) *Bitset {
	return &Bitset{length, make([]uint64, wordsNeeded(length))}
}

// Set bit i in the bitset to a 1
func (b *Bitset) Set(i uint64) {
	b.growBits(i)
	b.bits[i>>logWordSize] |= 1 << (i & (wordSize - 1))
}

func (b *Bitset) bitsLen() uint64 {
	return uint64(len(b.bits))
}

func (b *Bitset) growBits(i uint64) bool {
	if i < b.length {
		return false
	}
	needed := wordsNeeded(i + 1)
	bitsSize := b.bitsLen()
	if needed > bitsSize {
		newBits := make([]uint64, needed-bitsSize)
		b.bits = append(b.bits, newBits...)
	}
	b.length = i + 1
	return true
}
