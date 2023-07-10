package ecore

import (
	"fmt"
)

const valuesPerByte = 1 << 8 // 256

var indexModules = func() []int {
	xs := make([]int, 3*totalIndexes)
	for i := range xs {
		xs[i] = i % totalIndexes
	}
	return xs
}()

func errInvalidLetter(letter byte) error {
	return fmt.Errorf("Invalid letter %#U", letter)
}

func errInvalidIndex(index int) error {
	return fmt.Errorf("Invalid index %d", index)
}

func errInvalidLetterByIndex(letter byte, index int) error {
	return fmt.Errorf("Invalid letter %#U by index %d", letter, index)
}

func indexIsValid(index int) bool {
	return (0 <= index) && (index < totalIndexes)
}

//------------------------------------------------------------------------------

type indexOrigin struct{}

func (indexOrigin) indexToLetter(index int) (letter byte, ok bool) {
	if indexIsValid(index) {
		letter = byte(index + 'A')
		if ('A' <= letter) && (letter <= 'Z') {
			return letter, true
		}
	}
	return 0, false
}

func (indexOrigin) letterToIndex(letter byte) (index int, ok bool) {
	if ('a' <= letter) && (letter <= 'z') {
		return int(letter - 'a'), true
	}
	if ('A' <= letter) && (letter <= 'Z') {
		return int(letter - 'A'), true
	}
	return 0, false
}

//------------------------------------------------------------------------------

type indexCoder struct {
	enc [totalIndexes]byte
	dec [valuesPerByte]byte
}

func newIndexCoder() *indexCoder {

	ori := indexOrigin{}

	var enc [totalIndexes]byte
	for i := range enc {
		index := i
		letter, ok := ori.indexToLetter(index)
		if ok {
			enc[i] = letter
		}
	}

	var dec [valuesPerByte]byte
	for i := range dec {
		letter := byte(i)
		index, ok := ori.letterToIndex(letter)
		if ok {
			dec[i] = (byte(index) << 1) | 1
		}
	}
	return &indexCoder{
		enc: enc,
		dec: dec,
	}
}

func (p *indexCoder) indexToLetter(index int) (letter byte, ok bool) {
	if (0 <= index) && (index < len(p.enc)) {
		letter = p.enc[index]
		return letter, true
	}
	return 0, false
}

func (p *indexCoder) letterToIndex(letter byte) (index int, ok bool) {
	if (p.dec[letter] & 1) == 1 {
		index = int(p.dec[letter] >> 1)
		return index, true
	}
	return 0, false
}

//------------------------------------------------------------------------------

var globalIndexCoder = newIndexCoder()

// Letter to index
func letterToIndex(letter byte) (index int, ok bool) {
	return globalIndexCoder.letterToIndex(letter)
}

// Index to letter
func indexToLetter(index int) (letter byte, ok bool) {
	return globalIndexCoder.indexToLetter(index)
}

func LetterToIndex(letter byte) (index int, err error) {
	index, ok := letterToIndex(letter)
	if !ok {
		return 0, errInvalidLetter(letter)
	}
	return index, nil
}

func IndexToLetter(index int) (letter byte, err error) {
	letter, ok := indexToLetter(index)
	if !ok {
		return 0, errInvalidIndex(index)
	}
	return letter, nil
}

//------------------------------------------------------------------------------

func ParseIndexes(s string) ([]int, error) {
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

func parseIndexesN(s string, bs []int) error {
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
