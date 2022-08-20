package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gitchander/crypto/enigma/base26"
)

func main() {
	// testBase26Decode()
	testBase26EncodeHex()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testBase26Decode() {
	s := "NPFOMNOJBTMJGFGAPMDXCOJJZOJIMSKKMLJNMGDPGAMFDXDWH"
	bs, err := base26.DecodeString(s)
	checkError(err)
	fmt.Printf("%x\n", bs)
}

func testBase26EncodeHex() {

	//s := "0000000000000000"
	//s := "8888888888888888"
	s := "aaaaaaaa"
	//s := "ffffffff"
	fmt.Println("hex:", s)

	bs1, err := hex.DecodeString(s)
	checkError(err)

	enc := make([]byte, base26.EncodedLenMax(len(bs1)))
	n := base26.Encode(enc, bs1)
	enc = enc[:n]
	fmt.Printf("b26: %s\n", string(enc))

	bs2 := make([]byte, base26.DecodedLenMax(len(enc)))
	n, err = base26.Decode(bs2, enc)
	checkError(err)
	fmt.Printf("b26: %x\n", bs2[:n])
}
