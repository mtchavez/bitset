package bitset

import (
	"testing"
)

func BenchmarkSet(b *testing.B) {
	bs := New(1)
	var i uint64 = 0
	for ; i < uint64(b.N); i++ {
		bs.Set(i)
	}
}

func BenchmarkTest(b *testing.B) {
	bs := New(1)
	var i uint64 = 0
	for ; i < uint64(10000000); i++ {
		bs.Set(i)
	}
	i = 0
	for ; i < uint64(b.N); i++ {
		bs.Test(i)
	}
}
