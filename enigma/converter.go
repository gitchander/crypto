package enigma

import (
	"fmt"
)

var logConverters = true

type converter struct {
	name string
	cf   func(int) int
}

func newConverter(name string, cf func(int) int) *converter {
	return &converter{
		name: name,
		cf:   cf,
	}
}

func (c *converter) Do(index int) int {
	if logConverters {
		return c.doIndexLog(index)
	}
	return c.doIndex(index)
}

func (c *converter) doIndex(index int) int {
	return c.cf(index)
}

func (c *converter) doIndexLog(index int) int {

	inputLetter, _ := indexToLetter(index)
	index = c.doIndex(index)
	outputLetter, _ := indexToLetter(index)

	fmt.Printf("%-18s: %q -> %q\n", c.name, inputLetter, outputLetter)

	return index
}
