package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gitchander/crypto/magma"
)

var (
	key = []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	}

	syn = []byte{0xF1, 0x09, 0xAC, 0x11, 0x73, 0xB8, 0x04, 0x13}
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// type cipherSample struct {
// 	id         string
// 	config     Config
// 	key        []string
// 	plaintext  []string
// 	ciphertext []string
// }

func exampleBlock1() {

	fmt.Printf("key: [ % x ]\n", key)

	b, err := magma.NewCipher(key)
	checkError(err)

	data1 := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	data2 := make([]byte, b.BlockSize())
	data3 := make([]byte, b.BlockSize())

	b.Encrypt(data2, data1)
	fmt.Printf("[ % x ]\n", data2)

	b.Decrypt(data3, data2)
	fmt.Printf("[ % x ]\n", data3)
}

func exampleBlock2() {

	key, err := hex.DecodeString("FEDCBA9876543210FEDCBA9876543210FEDCBA9876543210FEDCBA9876543210")
	checkError(err)

	fmt.Printf("key: [ % x ]\n", key)

	b, err := magma.NewCipher(key)
	checkError(err)

	plaintext, err := hex.DecodeString("0123456789ABCDEF")
	checkError(err)

	ciphertext := make([]byte, b.BlockSize())
	plaintextDec := make([]byte, b.BlockSize())

	b.Encrypt(ciphertext, plaintext)
	fmt.Printf("ciphertext: [ % x ]\n", ciphertext)

	b.Decrypt(plaintextDec, ciphertext)
	fmt.Printf("plaintextDec: [ % x ]\n", plaintextDec)
}

func exampleBlock3() {

	rt := magma.RT2

	key := []byte{
		0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88,
		0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
		0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8,
		0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xd0,
	}

	fmt.Printf("key: [ % x ]\n", key)

	b, err := magma.NewCipherRT(rt, key)
	checkError(err)

	plaintext := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	checkError(err)

	ciphertext := make([]byte, b.BlockSize())
	plaintextDec := make([]byte, b.BlockSize())

	b.Encrypt(ciphertext, plaintext)
	fmt.Printf("ciphertext: [ %x ]\n", ciphertext)

	b.Decrypt(plaintextDec, ciphertext)
	fmt.Printf("plaintextDec: [ %x ]\n", plaintextDec)
}

func exampleStream() {

	b, err := magma.NewCipher(key)
	checkError(err)

	b1 := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	b2 := make([]byte, len(b1))
	b3 := make([]byte, len(b1))

	// Encrypt
	{
		se, err := magma.NewStreamCipher(b, syn)
		checkError(err)

		se.XORKeyStream(b2, b1)
	}

	// Decrypt
	{
		sd, err := magma.NewStreamCipher(b, syn)
		checkError(err)

		sd.XORKeyStream(b3[:5], b2[:5])
		sd.XORKeyStream(b3[5:9], b2[5:9])
		sd.XORKeyStream(b3[9:17], b2[9:17])
		sd.XORKeyStream(b3[17:], b2[17:])
	}

	const format = "%s: [ % x ]\n"
	fmt.Printf(format, "b1", b1)
	fmt.Printf(format, "b2", b2)
	fmt.Printf(format, "b3", b3)

	fmt.Println(string(b3))
}

func exampleCFB() {

	b, err := magma.NewCipher(key)
	checkError(err)

	data1 := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	data2 := make([]byte, len(data1))
	data3 := make([]byte, len(data1))

	// Encrypt
	{
		e := magma.NewCFBEncrypter(b, syn)

		e.XORKeyStream(data2, data1)
	}

	// Decrypt
	{
		d := magma.NewCFBDecrypter(b, syn)

		d.XORKeyStream(data3[:5], data2[:5])
		d.XORKeyStream(data3[5:9], data2[5:9])
		d.XORKeyStream(data3[9:17], data2[9:17])
		d.XORKeyStream(data3[17:], data2[17:])
	}

	const format = "%s: [ % x ]\n"

	fmt.Printf(format, "b1", data1)
	fmt.Printf(format, "b2", data2)
	fmt.Printf(format, "b3", data3)

	fmt.Println(string(data3))
}

func exampleHash() {

	b, err := magma.NewCipher(key)
	checkError(err)

	hash := magma.NewHash(b)

	data := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	hash.Write(data)
	fmt.Printf("hash: [ % x ]\n", hash.Sum(nil))

	s1 := "Hash is the common interface implemented by all hash functions"
	s2 := "Hash is the Common interface implemented by all hash functions"

	const format = "\"%s\": hash -> [ % x ]\n"

	hash.Reset()
	hash.Write([]byte(s1))
	fmt.Printf(format, s1, hash.Sum(nil))

	hash.Reset()
	hash.Write([]byte(s2))
	fmt.Printf(format, s2, hash.Sum(nil))
}

func main() {
	exampleBlock1()
	exampleBlock2()
	exampleBlock3()
	exampleStream()
	exampleCFB()
	exampleHash()
}
