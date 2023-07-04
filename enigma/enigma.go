package enigma

import (
	"fmt"
)

type Config struct {
	Plugboard string       `json:"plugboard,omitempty"`
	Rotors    RotorsConfig `json:"rotors"`
	Reflector string       `json:"reflector"`
}

type Enigma struct {
	plugboard *Plugboard
	rc        *rotorsCore
	reflector *Reflector
}

func New(c Config) (*Enigma, error) {

	plugboard, err := NewPlugboard(c.Plugboard)
	if err != nil {
		return nil, err
	}

	rc, err := newRotorsCore(c.Rotors)
	if err != nil {
		return nil, err
	}

	reflector, err := NewReflectorByID(c.Reflector)
	if err != nil {
		return nil, err
	}

	e := &Enigma{
		plugboard: plugboard,
		rc:        rc,
		reflector: reflector,
	}
	return e, nil
}

func (e *Enigma) Reset() {
	e.rc.reset()
}

func (e *Enigma) feed(index int) int {

	e.rc.rotorsRotate()

	index = e.plugboard.doForward(index)
	index = e.rc.rotorsForward(index)
	index = e.reflector.do(index)
	index = e.rc.rotorsBackward(index)
	index = e.plugboard.doBackward(index)

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
	outLetter, ok := indexToLetter(index)
	if !ok {
		panic(fmt.Errorf("invalid index %d", index))
	}

	return outLetter
}

func (e *Enigma) FeedString(s string) string {
	var (
		as = []byte(s)
		bs = make([]byte, len(as))
	)
	for i, a := range as {
		bs[i] = e.FeedLetter(a)
	}
	return string(bs)
}
