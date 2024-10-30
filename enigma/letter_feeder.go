package enigma

type LetterFeeder interface {
	FeedLetter(letter byte) (byte, error)
}
