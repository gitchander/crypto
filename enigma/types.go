package enigma

// lettersPerAlphabet
// nodesPerRotor
const nodes = 26

type convertTable [nodes]int

type dirRev struct {
	direct  convertTable
	reverse convertTable
}
