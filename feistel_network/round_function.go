package main

// round function
type RoundFunc func(k, r Word) Word

func RoundFuncXOR(k, r Word) Word {
	return xor(k, r)
}
