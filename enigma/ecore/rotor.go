package ecore

import (
	"fmt"
)

func errInvalidIndex(name string, index int) error {
	return fmt.Errorf("invalid (%s) value: have %d, want [%d:%d]", name, index, 0, positions)
}

type RotorConfig struct {
	Wiring    string
	Turnovers string
}

type Rotor struct {
	ct        coupleTable
	turnovers []bool
	ring      int
	position  int // position index
}

func NewRotor(rc RotorConfig) (*Rotor, error) {

	ct, err := parseWiring(rc.Wiring)
	if err != nil {
		return nil, err
	}

	turnovers, err := parseTurnovers(rc.Turnovers)
	if err != nil {
		return nil, err
	}

	r := &Rotor{
		ct:        ct,
		turnovers: turnovers,
	}
	return r, nil
}

func parseTurnovers(s string) ([]bool, error) {
	tis, err := ParseIndexes(s)
	if err != nil {
		return nil, err
	}
	turnovers := make([]bool, positions)
	for _, ti := range tis {
		turnovers[ti] = true
	}
	return turnovers, nil
}

func (r *Rotor) rotate() {
	r.position = (r.position + 1) % positions
}

func (r *Rotor) hasTurnover() bool {
	return r.turnovers[r.position]
}

func (r *Rotor) GetPosition() int {
	return r.position
}

func (r *Rotor) SetPosition(position int) error {
	if (0 <= position) && (position < positions) {
		r.position = position
		return nil
	}
	return errInvalidIndex("position", position)
}

func (r *Rotor) GetRing() int {
	return r.ring
}

func (r *Rotor) SetRing(ring int) error {
	if (0 <= ring) && (ring < positions) {
		r.ring = ring
		return nil
	}
	return errInvalidIndex("ring", ring)
}

func (r *Rotor) doV1(index int, ct *convertTable) int {
	index = mod((index - r.ring + r.position), positions)
	index = ct[index]
	index = mod((index + r.ring - r.position), positions)
	return index
}

func (r *Rotor) doV2(index int, ct *convertTable) int {
	index = (index - r.ring + r.position + positions) % positions
	index = ct[index]
	index = (index + r.ring - r.position + positions) % positions
	return index
}

// func (r *Rotor) doV3(index int, ct *convertTable) int {
// 	index = indexModules[(index - r.ring + r.position + positions)]
// 	index = ct[index]
// 	index = indexModules[(index + r.ring - r.position + positions)]
// 	return index
// }

func (r *Rotor) doTable(index int, ct *convertTable) int {
	//return r.doV1(index, ct)
	return r.doV2(index, ct)
	//return r.doV3(index, ct)
}

func (r *Rotor) DoForward(index int) int {
	return r.doTable(index, &(r.ct.forwardTable))
}

func (r *Rotor) DoBackward(index int) int {
	return r.doTable(index, &(r.ct.backwardTable))
}

func RotateRotors(rs []*Rotor) {
	hasPrev := true // Turnover last rotor.
	for i := len(rs) - 1; i >= 0; i-- {
		r := rs[i]
		ok := r.hasTurnover()
		if hasPrev || ok {
			r.rotate()
		}
		hasPrev = ok
	}
}
