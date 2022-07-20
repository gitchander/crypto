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

// example: "AV BS CG DL FU HZ IN KM OW RX"
func parsePlugboard(s string) (*Plugboard, error) {
	var dr dirRev
	for i := 0; i < nodes; i++ {
		dr.direct[i] = i
		dr.reverse[i] = i
	}
	if s == "" {
		return &Plugboard{dr}, nil
	}
	vs := strings.Split(s, " ")
	xs := make([]int, 2)
	for _, pair := range vs {
		err := parseLettersN(pair, xs)
		if err != nil {
			return nil, fmt.Errorf("invalid plugboard letters: %s", err)
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
