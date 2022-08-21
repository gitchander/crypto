package enigma

type RotorConfig struct {
	Wiring    string
	Turnovers string
}

type Rotor struct {
	dirRev
	turnovers []bool

	// Offset
	position int // position index

	ring int
}

func NewRotor(rc RotorConfig) (*Rotor, error) {

	dr, err := parseWiring(rc.Wiring)
	if err != nil {
		return nil, err
	}

	turnovers, err := parseTurnovers(rc.Turnovers)
	if err != nil {
		return nil, err
	}

	r := &Rotor{
		dirRev:    dr,
		turnovers: turnovers,
	}
	return r, nil
}

func parseTurnovers(s string) ([]bool, error) {
	tis, err := parseLetters(s)
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
	r.position = mod((r.position + 1), positions)
}

func (r *Rotor) hasTurnover() bool {
	return r.turnovers[r.position]
}

func (r *Rotor) getPosition() int {
	return r.position
}

func (r *Rotor) setPosition(position int) {
	r.position = mod(position, positions)
}

func (r *Rotor) getRing() int {
	return r.ring
}

func (r *Rotor) setRing(ring int) {
	r.ring = mod(ring, positions)
}

func (r *Rotor) do(index int, reverse bool) int {
	//return r.doV1(index, reverse)
	return r.doV2(index, reverse)
}

func (r *Rotor) doV1(index int, reverse bool) int {
	index = mod((index - r.ring + r.position), positions)
	if reverse {
		index = r.reverse[index]
	} else {
		index = r.direct[index]
	}
	index = mod((index + r.ring - r.position), positions)
	return index
}

func (r *Rotor) doV2(index int, reverse bool) int {
	index = indexModules[(index - r.ring + r.position + positions)]
	if reverse {
		index = r.reverse[index]
	} else {
		index = r.direct[index]
	}
	index = indexModules[(index + r.ring - r.position + positions)]
	return index
}

func (r *Rotor) doDirect(index int) int {
	return r.do(index, false)
}

func (r *Rotor) doReverse(index int) int {
	return r.do(index, true)
}

func rotateRotors(rs []*Rotor) {
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
