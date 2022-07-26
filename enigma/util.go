package enigma

import (
	"fmt"
)

var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var indexModules = func() []int {
	xs := make([]int, 3*positions)
	for i := range xs {
		xs[i] = i % positions
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

// var ErrInvalidWiring = errors.New("invalid wiring")

// mapping, wiring
func parseWiring(wiring string) (dirRev, error) {
	var dr dirRev
	bs := []byte(wiring)
	if len(bs) != positions {
		return dr, fmt.Errorf("wiring has invalid length %d", len(bs))
	}
	cs := make([]int, positions)
	for i, b := range bs {
		j, ok := letterToIndex(b)
		if !ok {
			return dr, fmt.Errorf("wiring has invalid letter %#U", b)
		}
		dr.direct[i] = j
		dr.reverse[j] = i
		cs[j]++
	}
	for i, c := range cs {
		if c == 0 {
			letter, _ := indexToLetter(i)
			return dr, fmt.Errorf("wiring has not letter %q", letter)
		} else if c > 1 {
			letter, _ := indexToLetter(i)
			return dr, fmt.Errorf("wiring has more than one letter %q", letter)
		}
	}
	return dr, nil
}

func parseDirectReverse(input, output string) (dirRev, error) {
	var dr dirRev
	var (
		as = []byte(input)
		bs = []byte(output)
	)
	if (len(as) != positions) || (len(bs) != positions) {
		return dr, fmt.Errorf("invalid rotor config")
	}
	for i := 0; i < positions; i++ {
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
