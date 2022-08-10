package main

import (
	"fmt"
	"log"

	"github.com/gitchander/crypto/enigma/base16"
)

func main() {
	testBase16Bytes()
	testBase16String()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testBase16Bytes() {
	//text1 := "Hello, World!"
	text1 := "Привет, Мир!"
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
	text1 := "Привет, Мир!"
	src := []byte(text1)
	s := base16.EncodeToString(src)
	fmt.Println(s)
	res, err := base16.DecodeString(s)
	checkError(err)
	text2 := string(res)
	fmt.Println(text2)
}
