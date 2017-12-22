package knot_hash

import "fmt"

type CircularArray struct {
	Elements *[]byte
	Size     int
}

func(ca *CircularArray) get(index int) byte {
	return (*ca.Elements)[index % ca.Size]
}

func (ca *CircularArray) set(index int, value byte) {
	(*ca.Elements)[index % ca.Size] = value
}

func(ca *CircularArray) Reverse(position int, length int) {
	if length > 1 {
		last := position + length - 1
		for i := 0; i < length / 2; i++ {
			a, b := ca.get(position+i), ca.get(last-i)
			ca.set(position+i, b)
			ca.set(last-i, a)
		}
	}
}

type KnotHash struct {
	elements *CircularArray
	lengths  []int
	position int
	skip     int
}

func newKnotHashFromLengths(lengths []int, size int) *KnotHash {
	elements := make([]byte, size)
	for i := 0; i < size; i++ {
		elements[i] = byte(i)
	}

	return &KnotHash{elements: &CircularArray{Elements: &elements, Size: size}, lengths: lengths, position: 0, skip: 0}
}

func newKnotHash(input string) *KnotHash {
	lengths := make([]int, len(input))
	for i, c := range input {
		lengths[i] = int(c)
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	return newKnotHashFromLengths(lengths, 256)
}

func(kh *KnotHash) nextIteration() {
	for _, length := range kh.lengths {
		kh.elements.Reverse(kh.position, length)
		kh.position = (kh.position + kh.skip + length) % kh.elements.Size
		kh.skip = kh.skip + 1
	}
}

func(kh *KnotHash) Rounds(totalRounds int) {
	for round := 0; round < totalRounds; round++ {
		kh.nextIteration()
	}
}

func(kh *KnotHash) String() string {
	representation := ""
	for i := 0; i < kh.elements.Size; i += 16 {
		value := (*kh.elements.Elements)[i]
		for j := i + 1; j < i + 16; j++ {
			value = value ^ (*kh.elements.Elements)[j]
		}
		representation += fmt.Sprintf("%02x", value)
	}
	return representation
}

func CalculateKnotHash(input string) string {
	kh := newKnotHash(input)
	kh.Rounds(64)

	return kh.String()
}