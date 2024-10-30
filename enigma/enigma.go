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
	plugboard  *ecore.Plugboard
	rotorBlock *ecore.RotorBlock
	reflector  *ecore.Reflector
}

var _ LetterFeeder = &Enigma{}

func New(c Config) (*Enigma, error) {

	plugboard, err := ecore.NewPlugboard(c.Plugboard)
	if err != nil {
		return nil, err
	}

	rotorBlock, err := newRotorBlock(c.Rotors)
	if err != nil {
		return nil, err
	}

	reflector, err := newReflectorByID(c.Reflector)
	if err != nil {
		return nil, err
	}

	e := &Enigma{
		plugboard:  plugboard,
		rotorBlock: rotorBlock,
		reflector:  reflector,
	}
	return e, nil
}

func (e *Enigma) feed(index int) int {

	e.rotorBlock.Rotate()

	index = e.plugboard.Forward(index)
	index = e.rotorBlock.Forward(index)
	index = e.reflector.Do(index)
	index = e.rotorBlock.Backward(index)
	index = e.plugboard.Backward(index)

	return index
}

func (e *Enigma) Reset() {
	e.rotorBlock.Reset()
}

// Encode / Decode
// Encrypt / Decrypt

func (e *Enigma) FeedLetter(letter byte) (byte, error) {

	index, err := ecore.LetterToIndex(letter)
	if err != nil {
		return 0, err
	}

	index = e.feed(index)

	outLetter, err := ecore.IndexToLetter(index)
	if err != nil {
		return 0, err
	}

	return outLetter, nil
}
