package enigma

// lettersPerAlphabet
// nodesPerRotor
// electrical contacts
const positions = 26

const (
	bitsPerByte   = 8
	valuesPerByte = 1 << bitsPerByte
)

type convertTable [positions]int

type dirRev struct {
	direct  convertTable
	reverse convertTable
}
