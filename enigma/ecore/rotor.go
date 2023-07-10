package ecore

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

func (r *Rotor) rotate() {
	r.position = (r.position + 1) % totalIndexes
}

func (r *Rotor) hasTurnover() bool {
	return r.turnovers[r.position]
}

func (r *Rotor) GetPosition() int {
	return r.position
}

func (r *Rotor) SetPosition(position int) error {
	if indexIsValid(position) {
		r.position = position
		return nil
	}
	return errInvalidIndex(position)
}

func (r *Rotor) GetRing() int {
	return r.ring
}

func (r *Rotor) SetRing(ring int) error {
	if indexIsValid(ring) {
		r.ring = ring
		return nil
	}
	return errInvalidIndex(ring)
}

func (r *Rotor) doV1(index int, ct *convertTable) int {
	index = mod((index - r.ring + r.position), totalIndexes)
	index = ct[index]
	index = mod((index + r.ring - r.position), totalIndexes)
	return index
}

func (r *Rotor) doV2(index int, ct *convertTable) int {
	index = (index - r.ring + r.position + totalIndexes) % totalIndexes
	index = ct[index]
	index = (index + r.ring - r.position + totalIndexes) % totalIndexes
	return index
}

func (r *Rotor) doV3(index int, ct *convertTable) int {
	index = indexModules[(index - r.ring + r.position + totalIndexes)]
	index = ct[index]
	index = indexModules[(index + r.ring - r.position + totalIndexes)]
	return index
}

func (r *Rotor) doTable(index int, ct *convertTable) int {
	//return r.doV1(index, ct)
	//return r.doV2(index, ct)
	return r.doV3(index, ct)
}

func (r *Rotor) Forward(index int) int {
	return r.doTable(index, &(r.ct.forwardTable))
}

func (r *Rotor) Backward(index int) int {
	return r.doTable(index, &(r.ct.backwardTable))
}
