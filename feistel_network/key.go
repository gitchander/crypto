package main

import (
	"errors"
)

var ErrorKeyLen = errors.New("wrong key len")

func expandKey(key []byte, we *wordEncoder) ([]Word, error) {
	n, rem := quoRem(len(key), sizeOfWord)
	if rem != 0 {
		return nil, ErrorKeyLen
	}
	ks := make([]Word, n)
	for i := range ks {
		ks[i] = we.getWord(key[i*sizeOfWord:])
	}
	return ks, nil
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
