package enigma

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

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

func JoinStrings(ss ...string) string {
	var b strings.Builder
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

func JoinLines(linePrefix string, lines ...string) string {
	var b strings.Builder
	for _, line := range lines {
		b.WriteString(linePrefix)
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return b.String()
}

func OnlyLetters(s string) string {
	rs := []rune(s)
	bs := make([]byte, 0, len(rs))
	for _, r := range rs {
		b, ok := runeSingleByte(r)
		if ok {
			index, ok := letterToIndex(b)
			if ok {
				letter, _ := indexToLetter(index)
				bs = append(bs, letter)
			}
		}
	}
	return string(bs)
}

func LinesToText(lines ...string) string {
	s := JoinLines("", lines...)
	s = OnlyLetters(s)
	return strings.ToUpper(s)
}

// one byte represent
func runeSingleByte(r rune) (byte, bool) {
	if uint32(r) < utf8.RuneSelf {
		return byte(r), true
	}
	return 0, false
}
