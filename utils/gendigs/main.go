package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var c Config
	flag.IntVar(&(c.Digits), "digs", 3, "number of digits")
	flag.StringVar(&(c.Filename), "filename", "", "result file name")
	flag.Parse()
	checkError(gen(c))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Digits   int
	Filename string
}

const (
	kilobyte = 1000
	megabyte = 1000 * 1000
)

func gen(c Config) error {
	filename := c.Filename
	if filename == "" {
		filename = fmt.Sprintf("digits-%02d.txt", c.Digits)
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriterSize(f, megabyte)
	defer w.Flush()

	n := c.Digits

	line := make([]byte, n+1) // digits + '\n'
	line[n] = '\n'

	digits := line[:n]
	fillBytes(digits, '0')

	for {
		_, err := w.Write(line)
		if err != nil {
			return err
		}
		if !(incDigits(digits)) {
			break
		}
	}

	return nil
}

func fillBytes(bs []byte, b byte) {
	for i := range bs {
		bs[i] = b
	}
}

var incTable = [256]byte{
	'0': '1',
	'1': '2',
	'2': '3',
	'3': '4',
	'4': '5',
	'5': '6',
	'6': '7',
	'7': '8',
	'8': '9',
	'9': '0',
}

func incDigits(digits []byte) bool {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = incTable[digits[i]]
		if digits[i] != '0' {
			return true
		}
	}
	return false
}
