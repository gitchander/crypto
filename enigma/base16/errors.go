package base16

import (
	"errors"
	"fmt"
)

var ErrLength = errors.New("enigma/base16: odd length base16 string")

type InvalidByteError byte

func (e InvalidByteError) Error() string {
	return fmt.Sprintf("enigma/base16: invalid byte: %#U", rune(e))
}
