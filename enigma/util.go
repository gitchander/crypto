package enigma

import (
	"strings"
	"unicode/utf8"

	"github.com/gitchander/crypto/enigma/ecore"
)

// one byte represent
func runeSingleByte(r rune) (byte, bool) {
	if uint32(r) < utf8.RuneSelf {
		return byte(r), true
	}
	return 0, false
}

//------------------------------------------------------------------------------

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
	var (
		rs = []rune(s)
		bs = make([]byte, 0, len(rs))
	)
	for _, r := range rs {
		b, ok := runeSingleByte(r)
		if ok {
			index, err := ecore.LetterToIndex(b)
			if err == nil {
				letter, _ := ecore.IndexToLetter(index)
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
