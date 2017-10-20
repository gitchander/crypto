package crygo

import (
	"errors"
	"fmt"
)

func newError(m string) error {
	return errors.New(fmt.Sprintf("crygo.%s", m))
}

func newErrorf(format string, a ...interface{}) error {
	return newError(fmt.Sprintf(format, a...))
}

var (
	ErrorSynLen = newError("wrong syn len")
	ErrorKeyLen = newError("wrong key len")

	ErrorTableLen    = newError("wrong table len")
	ErrorTableStruct = newError("wrong table struct")
)
