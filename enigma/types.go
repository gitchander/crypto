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

// coupleTable - forward and backward table.

type coupleTable struct {
	forwardTable  convertTable
	backwardTable convertTable
}
