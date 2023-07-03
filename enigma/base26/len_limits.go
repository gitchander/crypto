package base26

type limits struct{}

func (limits) EncodedLenMin(x int) int { return x * 5 / 4 } // x * 1.25
func (limits) EncodedLenMax(x int) int { return x * 2 }     // x * 2.0

func (limits) DecodedLenMin(x int) int { return x / 2 }     // x * 0.5
func (limits) DecodedLenMax(x int) int { return x * 4 / 5 } // x * 0.8

// EncodedLen
func EncodedLenMax(x int) int { return x * 2 }

// DecodedLen
func DecodedLenMax(x int) int { return x * 4 / 5 }
