package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func main() {

	text := `1947C0FCF2C8D64CB256FB722DA8E241`

	data, err := parseHex(text)
	checkError(err)

	//fmt.Printf("data: %X\n", data)

	ws, err := byteToUint64Slice(data)
	checkError(err)

	var b strings.Builder
	for i, w := range ws {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "0x%016x", w)
	}
	fmt.Println(b.String())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseHex(s string) ([]byte, error) {
	var (
		as = []byte(s)
		bs = make([]byte, 0, len(as))
	)
	for _, a := range as {
		if !byteIsSpace(a) {
			bs = append(bs, a)
		}
	}
	return hex.DecodeString(string(bs))
}

func byteIsSpace(b byte) bool {
	switch b {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}

func byteToUint64Slice(data []byte) ([]uint64, error) {

	const sizeOfUint64 = 8

	n, rem := quoRem(len(data), sizeOfUint64)

	if rem != 0 {
		return nil, fmt.Errorf("invalid data size %d", len(data))
	}

	ws := make([]uint64, n)
	for i := range ws {
		ws[i] = binary.LittleEndian.Uint64(data)
		data = data[sizeOfUint64:]
	}

	return ws, nil
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
