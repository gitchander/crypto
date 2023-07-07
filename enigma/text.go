package enigma

import (
	"strings"

	"github.com/gitchander/crypto/enigma/ecore"
)

type TextWorker struct{}

func (TextWorker) FeedTextPanicForeign(e *Enigma, text string) string {
	var (
		as = []byte(text)
		bs = make([]byte, len(as))
	)
	for i, a := range as {
		bs[i] = e.FeedLetter(a)
	}
	return string(bs)
}

// Include foreign chars
func (TextWorker) FeedTextIncludeForeign(e *Enigma, text string) string {
	var b strings.Builder
	for _, r := range text {
		x, ok := runeSingleByte(r)
		if ok {
			index, err := ecore.LetterToIndex(x)
			if err == nil {
				index = e.feed(index)
				x, _ = ecore.IndexToLetter(index)
			}
			b.WriteByte(x)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func (TextWorker) FeedTextIgnoreForeign(e *Enigma, text string) string {
	var b strings.Builder
	var i int
	for _, r := range text {
		x, ok := runeSingleByte(r)
		if ok {
			index, err := ecore.LetterToIndex(x)
			if err == nil {
				if (i > 0) && ((i % 5) == 0) {
					b.WriteByte(' ')
				}
				index = e.feed(index)
				x, _ = ecore.IndexToLetter(index)
				b.WriteByte(x)
				i++
			}
		}
	}
	return b.String()
}
