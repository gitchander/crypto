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

func (r *Rotor) doV3(index int, ct *convertTable) int {
	index = indexModules[(index - r.ring + r.position + positions)]
	index = ct[index]
	index = indexModules[(index + r.ring - r.position + positions)]
	return index
}

func (r *Rotor) doTable(index int, ct *convertTable) int {
	//return r.doV1(index, ct)
	return r.doV2(index, ct)
	//return r.doV3(index, ct)
}

func (r *Rotor) doForward(index int) int {
	return r.doTable(index, &(r.ct.forwardTable))
}

func (r *Rotor) doBackward(index int) int {
	return r.doTable(index, &(r.ct.backwardTable))
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
