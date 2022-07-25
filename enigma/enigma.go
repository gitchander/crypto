package enigma

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Config struct {
	Plugboard   string      `json:"plugboard,omitempty"`
	Rotors      []RotorInfo `json:"rotors"`
	ReflectorID string      `json:"reflector-id"`
}

type Enigma struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector
}

func New(c Config) (*Enigma, error) {

	plugboard, err := parsePlugboard(c.Plugboard)
	if err != nil {
		return nil, err
	}

	rotors := make([]*Rotor, len(c.Rotors))
	for i, ri := range c.Rotors {
		r, err := NewRotorByInfo(ri)
		if err != nil {
			return nil, err
		}
		rotors[i] = r
	}

	reflector, err := NewReflectorByID(c.ReflectorID)
	if err != nil {
		return nil, err
	}

	e := &Enigma{
		plugboard: plugboard,
		rotors:    rotors,
		reflector: reflector,
	}
	return e, nil
}

func (e *Enigma) rotate() {
	rotateRotors(e.rotors)
}

func makeRotorName(i int, suffix string) string {
	return fmt.Sprintf("rotor[%d]-%s", i, suffix)
}

func (e *Enigma) feed(index int) int {

	e.rotate()

	index = e.plugboard.Direct(index)

	n := len(e.rotors)
	for i := n - 1; i >= 0; i-- {
		index = e.rotors[i].Direct(index)
	}

	index = e.reflector.Direct(index)

	for i := 0; i < n; i++ {
		index = e.rotors[i].Reverse(index)
	}

	index = e.plugboard.Reverse(index)

	return index
}

// Encode
// Encrypt / Decrypt
func (e *Enigma) FeedLetter(letter byte) byte {

	index, ok := letterToIndex(letter)
	if !ok {
		panic(errInvalidLetter(letter))
	}

	index = e.feed(index)
	letter, _ = indexToLetter(index)

	return letter
}

func (e *Enigma) FeedString(s string) string {
	as := []byte(s)
	bs := make([]byte, len(as))
	for i, a := range as {
		bs[i] = e.FeedLetter(a)
	}
	return string(bs)
}

// Include foreign chars
func (e *Enigma) FeedIncludeForeign(s string) string {
	var b strings.Builder
	for _, r := range s {
		x, ok := runeSingleByte(r)
		if ok {
			index, ok := letterToIndex(x)
			if ok {
				index = e.feed(index)
				x, _ = indexToLetter(index)
			}
			b.WriteByte(x)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func (e *Enigma) FeedIgnoreForeign(s string) string {
	var b strings.Builder
	var i int
	for _, r := range s {
		x, ok := runeSingleByte(r)
		if ok {
			index, ok := letterToIndex(x)
			if ok {
				if (i > 0) && ((i % 5) == 0) {
					b.WriteByte(' ')
				}
				index = e.feed(index)
				x, _ = indexToLetter(index)
				b.WriteByte(x)
				i++
			}
		}
	}
	return b.String()
}

// one byte represent
func runeSingleByte(r rune) (byte, bool) {
	if uint32(r) < utf8.RuneSelf {
		return byte(r), true
	}
	return 0, false
}
