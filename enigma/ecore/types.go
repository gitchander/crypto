package ecore

type convertTable [positions]int

// coupleTable - forward and backward table.

type coupleTable struct {
	forwardTable  convertTable
	backwardTable convertTable
}
