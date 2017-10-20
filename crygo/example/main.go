package main

import (
	"fmt"

	"github.com/gitchander/crypto/crygo"
)

var (
	table = crygo.NewTableDefault()

	key = []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	}

	syn = []byte{0xF1, 0x09, 0xAC, 0x11, 0x73, 0xB8, 0x04, 0x13}
)

func BlockExample() error {

	fmt.Printf("key: [ % x ]\n", key)

	blockCipher, err := crygo.NewBlockCipher(table, key)
	if err != nil {
		return err
	}

	b1 := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	b2 := make([]byte, crygo.BlockSize)
	b3 := make([]byte, crygo.BlockSize)

	blockCipher.Encrypt(b2, b1)
	fmt.Printf("[ % x ]\n", b2)

	blockCipher.Decrypt(b3, b2)
	fmt.Printf("[ % x ]\n", b3)

	return nil
}

func StreamExample() error {

	blockCipher, err := crygo.NewBlockCipher(table, key)
	if err != nil {
		return err
	}

	b1 := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	b2 := make([]byte, len(b1))
	b3 := make([]byte, len(b1))

	// Encrypt
	{
		se, err := crygo.NewStreamCipher(blockCipher, syn)
		if err != nil {
			return err
		}

		se.XORKeyStream(b2, b1)
	}

	// Decrypt
	{
		sd, err := crygo.NewStreamCipher(blockCipher, syn)
		if err != nil {
			return err
		}

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

	return nil
}

func CFBExample() error {

	blockCipher, err := crygo.NewBlockCipher(table, key)
	if err != nil {
		return err
	}

	b1 := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	b2 := make([]byte, len(b1))
	b3 := make([]byte, len(b1))

	// Encrypt
	{
		e := crygo.NewCFBEncrypter(blockCipher, syn)

		e.XORKeyStream(b2, b1)
	}

	// Decrypt
	{
		d := crygo.NewCFBDecrypter(blockCipher, syn)

		d.XORKeyStream(b3[:5], b2[:5])
		d.XORKeyStream(b3[5:9], b2[5:9])
		d.XORKeyStream(b3[9:17], b2[9:17])
		d.XORKeyStream(b3[17:], b2[17:])
	}

	const format = "%s: [ % x ]\n"

	fmt.Printf(format, "b1", b1)
	fmt.Printf(format, "b2", b2)
	fmt.Printf(format, "b3", b3)

	fmt.Println(string(b3))

	return nil
}

func HashExample() error {

	blockCipher, err := crygo.NewBlockCipher(table, key)
	if err != nil {
		return err
	}

	hash := crygo.NewHash(blockCipher)

	b := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	hash.Write(b)
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

	return nil
}

func main() {

	if err := BlockExample(); err != nil {
		fmt.Println(err.Error())
	}

	if err := StreamExample(); err != nil {
		fmt.Println(err.Error())
	}

	if err := CFBExample(); err != nil {
		fmt.Println(err.Error())
	}

	if err := HashExample(); err != nil {
		fmt.Println(err.Error())
	}
}
