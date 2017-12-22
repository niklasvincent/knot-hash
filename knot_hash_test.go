package knot_hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateKnotHashSimpleCase(t *testing.T) {
	lengths := []int{3, 4, 1, 5}
	kh := newKnotHashFromLengths(lengths, 5)
	kh.Rounds(1)

	assert.Equal(t, []byte{3, 4, 2, 1, 0}, *kh.elements.Elements)
}

func TestCalculateKnotHashPuzzleInput(t *testing.T) {
	lengths := []int{187, 254, 0, 81, 169, 219, 1, 190, 19, 102, 255, 56, 46, 32, 2, 216}
	kh := newKnotHashFromLengths(lengths, 256)
	kh.Rounds(1)

	assert.Equal(t, 1980, int((*kh.elements.Elements)[0]) * int((*kh.elements.Elements)[1]))
}

func TestCalculateKnotHashExampleStrings(t *testing.T) {
	assert.Equal(t, "a2582a3a0e66e6e86e3812dcb672a272", CalculateKnotHash(""))
	assert.Equal(t, "33efeb34ea91902bb2f59c9920caa6cd", CalculateKnotHash("AoC 2017"))
	assert.Equal(t, "3efbe78a8d82f29979031a4aa0b16a9d", CalculateKnotHash("1,2,3"))
	assert.Equal(t, "63960835bcdc130f0b66d7ff4f6a5a8e", CalculateKnotHash("1,2,4"))
	assert.Equal(t, "899124dac21012ebc32e2f4d11eaec55", CalculateKnotHash("187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216"))
}

func BenchmarkCalculateKnotHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CalculateKnotHash("187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216")
	}
}