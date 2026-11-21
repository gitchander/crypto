package main

import (
	"fmt"
	"log"

	"github.com/gitchander/crypto/enigma/base16"
)

func main() {
	// testBase16Bytes()
	// testBase16String()
	exampleBinary()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testBase16Bytes() {
	text1 := "Hello, World!"
	src := []byte(text1)
	dst := make([]byte, base16.EncodedLen(len(src)))
	base16.Encode(dst, src)
	res := make([]byte, base16.DecodedLen(len(dst)))
	_, err := base16.Decode(res, dst)
	checkError(err)
	text2 := string(res)
	fmt.Println(text2)
}

func testBase16String() {
	text1 := "Hello, World!"
	src := []byte(text1)
	s := base16.EncodeToString(src)
	fmt.Println(s)
	res, err := base16.DecodeString(s)
	checkError(err)
	text2 := string(res)
	fmt.Println(text2)
}

func exampleBinary() {
	src := []byte{0x10, 0x32, 0x54, 0x76, 0x98, 0xBA, 0xDC, 0xFE}
	dst := make([]byte, base16.EncodedLen(len(src)))
	base16.Encode(dst, src)
	text := string(dst)
	fmt.Println(text) // ABCDEFGHJKLMNPRS
	res := make([]byte, base16.DecodedLen(len(dst)))
	_, err := base16.Decode(res, dst)
	checkError(err)
	fmt.Printf("[% X]\n", res) // [10 32 54 76 98 BA DC FE]
}
