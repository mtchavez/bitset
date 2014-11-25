package bitset

import "math"

type Bitset struct {
	n    uint64
	bits []uint64
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
