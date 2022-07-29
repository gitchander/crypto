package enigma

import (
	"fmt"
	"strings"
)

var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var indexModules = func() []int {
	xs := make([]int, 3*positions)
	for i := range xs {
		xs[i] = i % positions
	}
	return xs
}()

// Letter to position
func letterToIndex(letter byte) (index int, ok bool) {
	if ('a' <= letter) && (letter <= 'z') {
		return int(letter - 'a'), true
	}
	if ('A' <= letter) && (letter <= 'Z') {
		return int(letter - 'A'), true
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

// mapping, wiring
func parseWiring(wiring string) (dirRev, error) {

	var dr dirRev
	err := ValidateWiring(wiring)
	if err != nil {
		return dr, err
	}

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

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func duplicateRunes(s string) (rs []rune) {
	m := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := m[r]; ok {
			rs = append(rs, r)
		} else {
			m[r] = struct{}{}
		}
	}
	return rs
}

func JoinLines(lines []string) string {
	var b strings.Builder
	for _, line := range lines {
		b.WriteString(line)
	}
	return b.String()
}

func OnlyLetters(s string) string {
	as := []byte(s)
	bs := make([]byte, 0, len(as))
	for _, a := range as {
		if byteIsLetter(a) {
			bs = append(bs, a)
		}
	}
	return string(bs)
}

func byteIsLetter(b byte) bool {
	if ('A' <= b) && (b <= 'Z') {
		return true
	}
	if ('a' <= b) && (b <= 'z') {
		return true
	}
	return false
}

func LinesToText(lines []string) string {
	s := JoinLines(lines)
	s = OnlyLetters(s)
	return strings.ToUpper(s)
}
