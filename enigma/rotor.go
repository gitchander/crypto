package enigma

type RotorConfig struct {
	Wiring    string
	Turnovers string
}

type Rotor struct {
	ct        coupleTable
	turnovers []bool

	// Offset
	position int // position index

	ring int
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

func (r *Rotor) doV1(index int, backward bool) int {
	index = mod((index - r.ring + r.position), positions)
	if backward {
		index = r.ct.backwardTable[index]
	} else {
		index = r.ct.forwardTable[index]
	}
	index = mod((index + r.ring - r.position), positions)
	return index
}

func (r *Rotor) doV2(index int, backward bool) int {
	index = indexModules[(index - r.ring + r.position + positions)]
	if backward {
		index = r.ct.backwardTable[index]
	} else {
		index = r.ct.forwardTable[index]
	}
	index = indexModules[(index + r.ring - r.position + positions)]
	return index
}

func (r *Rotor) do(index int, backward bool) int {
	//return r.doV1(index, backward)
	return r.doV2(index, backward)
}

func (r *Rotor) doForward(index int) int {
	return r.do(index, false)
}

func (r *Rotor) doBackward(index int) int {
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
