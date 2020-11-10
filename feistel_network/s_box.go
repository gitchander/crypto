package main

type SBox interface {
	Substitute(Word) Word
}
