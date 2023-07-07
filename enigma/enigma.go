package enigma

import (
	"github.com/gitchander/crypto/enigma/ecore"
)

type Config struct {
	Plugboard string       `json:"plugboard,omitempty"`
	Rotors    RotorsConfig `json:"rotors"`
	Reflector string       `json:"reflector"`
}

type Enigma struct {
	plugboard *ecore.Plugboard
	rc        *rotorsCore
	reflector *ecore.Reflector
}

func New(c Config) (*Enigma, error) {

	plugboard, err := ecore.NewPlugboard(c.Plugboard)
	if err != nil {
		return nil, err
	}

	rc, err := newRotorsCore(c.Rotors)
	if err != nil {
		return nil, err
	}

	reflector, err := reflectorByID(c.Reflector)
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

	e.rc.rotate()

	index = e.plugboard.DoForward(index)
	index = e.rc.doForward(index)
	index = e.reflector.Do(index)
	index = e.rc.doBackward(index)
	index = e.plugboard.DoBackward(index)

	return index
}

// Encode
// Encrypt / Decrypt
func (e *Enigma) FeedLetter(letter byte) byte {

	index, err := ecore.LetterToIndex(letter)
	if err != nil {
		panic(err)
	}

	index = e.feed(index)
	outLetter, err := ecore.IndexToLetter(index)
	if err != nil {
		panic(err)
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
