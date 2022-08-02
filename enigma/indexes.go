package enigma

import (
	"fmt"
)

const maskByteIsLetter = 1 << (bitsPerByte - 1) // last bit is set

var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var indexModules = func() []int {
	xs := make([]int, 3*positions)
	for i := range xs {
		xs[i] = i % positions
	}
	return xs
}()

var tableIndexes = func() (table [valuesPerByte]byte) {
	for i := range table {
		index, ok := letterToIndexSlow(byte(i))
		if ok {
			table[i] = maskByteIsLetter | byte(index)
		}
	}
	return table
}()

func letterToIndexSlow(letter byte) (index int, ok bool) {
	if ('a' <= letter) && (letter <= 'z') {
		return int(letter - 'a'), true
	}
	if ('A' <= letter) && (letter <= 'Z') {
		return int(letter - 'A'), true
	}
	return 0, false
}

//--------------------------------------
// bitwise operators:
// a | b  - bitwise OR
// a & b  - bitwise AND
// a ^ b  - bitwise XOR
// a << b - bitwise left shift
// a >> b - bitwise right shift
// a &^ b - bitwise AND NOT
//--------------------------------------
// Letter to index
func letterToIndex(letter byte) (index int, ok bool) {
	x := tableIndexes[letter]
	if (x & maskByteIsLetter) != 0 {
		return int(x ^ maskByteIsLetter), true
	}
	return 0, false
}

// Position to letter
func indexToLetter(index int) (letter byte, ok bool) {
	if (0 <= index) && (index < len(alphabet)) {
		return alphabet[index], true
	}
	return 0, false
}

func parseIndex(s string) (index int, err error) {
	if len(s) != 1 {
		return 0, fmt.Errorf("invalid letter %q", s)
	}
	letter := s[0]
	index, ok := letterToIndex(letter)
	if !ok {
		return 0, fmt.Errorf("invalid letter %#U", letter)
	}
	return index, nil
}
