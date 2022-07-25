package enigma

import (
	"errors"
	"fmt"
)

var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func indexMod(index int) int {
	//return index % nodes
	return indexModules_[index]
}

var indexModules_ = func() []int {
	xs := make([]int, 2*nodes) // [0..25,26..51]
	for i := range xs {
		xs[i] = i % nodes
	}
	return xs
}()

func letterToIndex(letter byte) (index int, ok bool) {
	if ('a' <= letter) && (letter <= 'z') {
		return int(letter - 'a'), true
	}
	if ('A' <= letter) && (letter <= 'Z') {
		return int(letter - 'A'), true
	}
	return 0, false
}

func indexToLetter(index int) (letter byte, ok bool) {
	if (0 <= index) && (index <= 25) {
		return byte(index + 'A'), true
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

func errInvalidLetter(letter byte) error {
	//return fmt.Errorf("invalid letter %#U", letter)
	return fmt.Errorf("there is an invalid letter %#U", letter)
}

func errInvalidLetterByIndex(letter byte, index int) error {
	return fmt.Errorf("invalid letter %#U by index %d", letter, index)
}

func parseLetters(s string) ([]int, error) {
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

// func DoLetter(letter byte, f func(int) int) byte {
// 	index, ok := letterToIndex(letter)
// 	if !ok {
// 		panic("invalid input letter")
// 	}
// 	out, ok := indexToLetter(f(index))
// 	if !ok {
// 		panic("invalid output letter")
// 	}
// 	return out
// }

var ErrInvalidWiring = errors.New("invalid wiring")

// mapping, wiring
func parseWiring(wiring string) (dirRev, error) {
	var dr dirRev
	bs := []byte(wiring)
	if len(bs) != nodes {
		return dr, ErrInvalidWiring
	}
	for i, b := range bs {
		j, ok := letterToIndex(b)
		if !ok {
			return dr, ErrInvalidWiring
		}
		dr.direct[i] = j
		dr.reverse[j] = i
	}
	return dr, nil
}

func parseDirectReverse(input, output string) (dirRev, error) {
	var dr dirRev
	var (
		as = []byte(input)
		bs = []byte(output)
	)
	if (len(as) != nodes) || (len(bs) != nodes) {
		return dr, fmt.Errorf("invalid rotor config")
	}
	for i := 0; i < nodes; i++ {
		ai, ok := letterToIndex(as[i])
		if !ok {
			return dr, fmt.Errorf("invalid %s letter %q by index %d", "input", as[i], i)
		}
		bi, ok := letterToIndex(bs[i])
		if !ok {
			return dr, fmt.Errorf("invalid %s letter %q by index %d", "output", bs[i], i)
		}
		dr.direct[ai] = bi
		dr.reverse[bi] = ai
	}
	return dr, nil
}

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

// func parseIndex(b byte) (int, error) {
// 	index, ok := letterToIndex(b)
// 	if !ok {
// 		return 0, errInvalidLetter(b)
// 	}
// 	return index, nil
// }
