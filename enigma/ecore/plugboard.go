package ecore

import (
	"fmt"
	"strings"
)

type Plugboard struct {
	coupleTable
}

func NewPlugboard(s string) (*Plugboard, error) {

	err := ValidatePlugboard(s)
	if err != nil {
		return nil, err
	}

	var ct coupleTable
	for i := 0; i < totalIndexes; i++ {
		ct.forwardTable[i] = i
		ct.backwardTable[i] = i
	}
	if s == "" {
		return &Plugboard{ct}, nil
	}

	uniqueMap := make(map[byte]struct{})

	vs := strings.Split(s, " ")
	xs := make([]int, 2)
	for _, pair := range vs {
		err := parseIndexesN(pair, xs)
		if err != nil {
			return nil, fmt.Errorf("invalid plugboard letters: %s", err)
		}

		// Check duplicates:
		for _, r := range pair {
			b := byte(r)
			if _, ok := uniqueMap[b]; ok {
				return nil, fmt.Errorf("plugboard has duplicate of char %q", b)
			}
			uniqueMap[b] = struct{}{}
		}

		var (
			a = xs[0]
			b = xs[1]
		)

		ct.forwardTable[a] = b
		ct.forwardTable[b] = a

		ct.backwardTable[a] = b
		ct.backwardTable[b] = a
	}
	return &Plugboard{ct}, nil
}

func (p *Plugboard) Forward(index int) int {
	return p.forwardTable[index]
}

func (p *Plugboard) Backward(index int) int {
	return p.backwardTable[index]
}
