package ecore

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

// --------------------------------------
// bitwise operators:
// a | b  - bitwise OR
// a & b  - bitwise AND
// a ^ b  - bitwise XOR
// a << b - bitwise left shift
// a >> b - bitwise right shift
// a &^ b - bitwise AND NOT
// --------------------------------------
// Letter to index
func letterToIndex(letter byte) (index int, ok bool) {
	x := tableIndexes[letter]
	if (x & maskByteIsLetter) != 0 {
		return int(x &^ maskByteIsLetter), true
	}
	return 0, false
}

func LetterToIndex(letter byte) (index int, err error) {
	index, ok := letterToIndex(letter)
	if !ok {
		return 0, errInvalidLetter(letter)
	}
	return index, nil
}

// Position to letter
func indexToLetter(index int) (letter byte, ok bool) {
	if (0 <= index) && (index < len(alphabet)) {
		return alphabet[index], true
	}
	return 0, false
}

func IndexToLetter(index int) (letter byte, err error) {
	letter, ok := indexToLetter(index)
	if !ok {
		return 0, fmt.Errorf("Invalid index %d", index)
	}
	return letter, nil
}

// fmt.Errorf("invalid index %d", index)

func ParseIndexes(s string) ([]int, error) {
	var (
		as = []byte(s)
		bs = make([]int, len(as))
	)
	for i, a := range as {
		b, ok := letterToIndex(a)
		if !ok {
			return nil, errInvalidLetterByIndex(a, i)
		}
		bs[i] = b
	}
	return bs, nil
}

func parseLettersN(s string, bs []int) error {
	as := []byte(s)
	if len(as) < len(bs) {
		return fmt.Errorf("insufficient length of letters: have %d, want %d", len(as), len(bs))
	}
	for i, a := range as {
		b, ok := letterToIndex(a)
		if !ok {
			return errInvalidLetterByIndex(a, i)
		}
		bs[i] = b
	}
	return nil
}
