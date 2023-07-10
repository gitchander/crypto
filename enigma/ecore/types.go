package ecore

const totalIndexes = 26

type convertTable [totalIndexes]int

// coupleTable - forward and backward table.

type coupleTable struct {
	forwardTable  convertTable
	backwardTable convertTable
}
