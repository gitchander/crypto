package ecore

type ReflectorConfig struct {
	Wiring string
}

type Reflector struct {
	ct coupleTable
}

func NewReflector(rc ReflectorConfig) (*Reflector, error) {
	ct, err := parseWiring(rc.Wiring)
	if err != nil {
		return nil, err
	}
	r := &Reflector{
		ct: ct,
	}
	return r, nil
}

func (r *Reflector) Do(index int) int {
	return r.ct.forwardTable[index]
}
