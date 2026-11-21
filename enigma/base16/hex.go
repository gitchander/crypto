package base16

import (
	"encoding/hex"
)

func exampleHex() {
	text1 := "Hello, World!"
	src := []byte(text1)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	res := make([]byte, hex.DecodedLen(len(dst)))
	hex.Decode(res, dst)
	text2 := string(res)
	_ = text2
}
