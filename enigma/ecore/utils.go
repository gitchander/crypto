package ecore

import (
	"fmt"
)

func errInvalidLetter(letter byte) error {
	return fmt.Errorf("Invalid letter %#U", letter)
}

func errInvalidLetterByIndex(letter byte, index int) error {
	return fmt.Errorf("Invalid letter %#U by index %d", letter, index)
}

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func findDuplicate[T comparable](as []T) (T, bool) {
	m := make(map[T]struct{})
	for _, a := range as {
		if _, ok := m[a]; ok {
			return a, true
		} else {
			m[a] = struct{}{}
		}
	}
	var zero T
	return zero, false
}

// mapping, wiring
func parseWiring(wiring string) (coupleTable, error) {

	var ct coupleTable
	err := ValidateWiring(wiring)
	if err != nil {
		return ct, err
	}

	bs := []byte(wiring)
	if len(bs) != positions {
		return ct, fmt.Errorf("wiring has invalid length %d", len(bs))
	}
	cs := make([]int, positions)
	for i, b := range bs {
		j, ok := letterToIndex(b)
		if !ok {
			return ct, fmt.Errorf("wiring has invalid letter %#U", b)
		}
		ct.forwardTable[i] = j
		ct.backwardTable[j] = i
		cs[j]++
	}
	for i, c := range cs {
		if c == 0 {
			letter, _ := indexToLetter(i)
			return ct, fmt.Errorf("wiring has not letter %q", letter)
		} else if c > 1 {
			letter, _ := indexToLetter(i)
			return ct, fmt.Errorf("wiring has more than one letter %q", letter)
		}
	}
	return ct, nil
}
