package enigma

import (
	"strings"

	"github.com/gitchander/crypto/enigma/ecore"
)

type TextFeeder interface {
	FeedText(text string) (string, error)
}

//------------------------------------------------------------------------------

type trueTextFeeder struct {
	e *Enigma
}

// RightTextFeeder
func TrueTextFeeder(e *Enigma) TextFeeder {
	return &trueTextFeeder{e}
}

func (p *trueTextFeeder) FeedText(text string) (string, error) {

	p.e.Reset()

	var (
		as = []byte(text)
		bs = make([]byte, len(as))
	)
	for i, a := range as {
		b, err := p.e.FeedLetter(a)
		if err != nil {
			return "", err
		}
		bs[i] = b
	}
	return string(bs), nil
}

//------------------------------------------------------------------------------

type includeForeign struct {
	e *Enigma
}

func IncludeForeign(e *Enigma) TextFeeder {
	return &includeForeign{e}
}

// Include foreign chars
func (p *includeForeign) FeedText(text string) (string, error) {

	p.e.Reset()

	var b strings.Builder
	for _, r := range text {
		x, ok := runeToSingleByte(r)
		if ok {
			index, err := ecore.LetterToIndex(x)
			if err == nil {
				index = p.e.feed(index)
				x, _ = ecore.IndexToLetter(index)
			}
			b.WriteByte(x)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String(), nil
}

//------------------------------------------------------------------------------

type ignoreForeign struct {
	e *Enigma
}

func IgnoreForeign(e *Enigma) TextFeeder {
	return &ignoreForeign{e}
}

func (p *ignoreForeign) FeedText(text string) (string, error) {

	p.e.Reset()

	var b strings.Builder
	for _, r := range text {
		x, ok := runeToSingleByte(r)
		if ok {
			index, err := ecore.LetterToIndex(x)
			if err == nil {
				index = p.e.feed(index)
				x, _ = ecore.IndexToLetter(index)
				b.WriteByte(x)
			}
		}
	}
	return b.String(), nil
}
