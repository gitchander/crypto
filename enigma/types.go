package enigma

// lettersPerAlphabet
// nodesPerRotor
// electrical contacts
const positions = 26

type convertTable [positions]int

type dirRev struct {
	direct  convertTable
	reverse convertTable
}
