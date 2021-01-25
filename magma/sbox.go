package magma

type SBox interface {
	Substitute(word) word
}
