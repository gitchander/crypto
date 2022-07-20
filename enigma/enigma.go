package enigma

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Config struct {
	// Plug board settings:
	Plugboard string // PO ML IU KJ NH YT GB VF RE DC

	Rotors []RotorInfo

	Reflector ReflectorConfig
}

type Enigma struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector

	cs []*converter
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

	reflector, err := NewReflector(c.Reflector)
	if err != nil {
		return nil, err
	}

	e := &Enigma{
		plugboard: plugboard,
		rotors:    rotors,
		reflector: reflector,
	}

	e.cs = e.makeConverters()

	return e, nil
}

func (e *Enigma) rotate() {
	rotateRotors(e.rotors)
}

func (e *Enigma) makeConverters() []*converter {

	var cs []*converter

	cs = append(cs, newConverter("plugboard-direct", e.plugboard.Direct))

	n := len(e.rotors)
	for i := n - 1; i >= 0; i-- {
		var (
			r         = e.rotors[i]
			rotorName = makeRotorName(i, "direct")
			c         = newConverter(rotorName, r.Direct)
		)
		cs = append(cs, c)
	}

	cs = append(cs, newConverter("reflector", e.reflector.Direct))

	for i := 0; i < n; i++ {
		var (
			r         = e.rotors[i]
			rotorName = makeRotorName(i, "reverse")
			c         = newConverter(rotorName, r.Reverse)
		)
		cs = append(cs, c)
	}

	cs = append(cs, newConverter("plugboard-reverse", e.plugboard.Reverse))

	return cs
}

func makeRotorName(i int, suffix string) string {
	return fmt.Sprintf("rotor[%d]-%s", i, suffix)
}

func (e *Enigma) doConverters(index int) int {
	for _, c := range e.cs {
		index = c.Do(index)
	}
	return index
}

func (e *Enigma) feed(index int) int {

	e.rotate()

	if false {
		return e.doConverters(index)
	}

	//--------------------------------------------------------------------------
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
