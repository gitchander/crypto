package enigma

import (
	"strings"

	"github.com/gitchander/crypto/enigma/ecore"
)

type TextFormatter struct {
	LettersPerGroup int
	GroupsPerLine   int
}

var DefaultTextFormatter = TextFormatter{
	LettersPerGroup: 4,
	GroupsPerLine:   12,
}

func (tf TextFormatter) FormatText(text string) string {
	var b strings.Builder
	var i int
	for _, r := range text {
		x, ok := runeToSingleByte(r)
		if ok {
			index, err := ecore.LetterToIndex(x)
			if err == nil {
				if i > 0 {
					if (i % (tf.GroupsPerLine * tf.LettersPerGroup)) == 0 {
						b.WriteByte('\n')
					} else if (i % tf.LettersPerGroup) == 0 {
						b.WriteByte(' ')
					}
				}
				x, _ = ecore.IndexToLetter(index)
				b.WriteByte(x)
				i++
			}
		}
	}
	return b.String()
}
