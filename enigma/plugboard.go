package enigma

import (
	"fmt"
	"strings"
)

type Plugboard struct {
	dirRev
}

func NewPlugboard(s string) (*Plugboard, error) {
	return parsePlugboard(s)
}

func parsePlugboard(s string) (*Plugboard, error) {

	err := ValidatePlugboard(s)
	if err != nil {
		return nil, err
	}

	var dr dirRev
	for i := 0; i < positions; i++ {
		dr.direct[i] = i
		dr.reverse[i] = i
	}
	if s == "" {
		return &Plugboard{dr}, nil
	}

	uniqueMap := make(map[byte]struct{})

	vs := strings.Split(s, " ")
	xs := make([]int, 2)
	for _, pair := range vs {
		err := parseLettersN(pair, xs)
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

		dr.direct[a] = b
		dr.direct[b] = a

		dr.reverse[a] = b
		dr.reverse[b] = a
	}
	return &Plugboard{dr}, nil
}

func (p *Plugboard) Direct(index int) int {
	return p.direct[index]
}

func (p *Plugboard) Reverse(index int) int {
	return p.reverse[index]
}
