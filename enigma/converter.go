package enigma

// import (
// 	"fmt"
// )

// var logConverters = true

// type converter struct {
// 	name string
// 	cf   func(int) int
// }

// func newConverter(name string, cf func(int) int) *converter {
// 	return &converter{
// 		name: name,
// 		cf:   cf,
// 	}
// }

// func (c *converter) Do(index int) int {
// 	if logConverters {
// 		return c.doIndexLog(index)
// 	}
// 	return c.doIndex(index)
// }

// func (c *converter) doIndex(index int) int {
// 	return c.cf(index)
// }

// func (c *converter) doIndexLog(index int) int {

// 	inputLetter, _ := indexToLetter(index)
// 	index = c.doIndex(index)
// 	outputLetter, _ := indexToLetter(index)

// 	fmt.Printf("%-18s: %q -> %q\n", c.name, inputLetter, outputLetter)

// 	return index
// }

// func makeConverters(e *Enigma) []*converter {

// 	var cs []*converter

// 	cs = append(cs, newConverter("plugboard-direct", e.plugboard.Direct))

// 	n := len(e.rotors)
// 	for i := n - 1; i >= 0; i-- {
// 		var (
// 			r         = e.rotors[i]
// 			rotorName = makeRotorName(i, "direct")
// 			c         = newConverter(rotorName, r.Direct)
// 		)
// 		cs = append(cs, c)
// 	}

// 	cs = append(cs, newConverter("reflector", e.reflector.Direct))

// 	for i := 0; i < n; i++ {
// 		var (
// 			r         = e.rotors[i]
// 			rotorName = makeRotorName(i, "reverse")
// 			c         = newConverter(rotorName, r.Reverse)
// 		)
// 		cs = append(cs, c)
// 	}

// 	cs = append(cs, newConverter("plugboard-reverse", e.plugboard.Reverse))

// 	return cs
// }

// func doConverters(cs []*converter, index int) int {
// 	for _, c := range cs {
// 		index = c.Do(index)
// 	}
// 	return index
// }
