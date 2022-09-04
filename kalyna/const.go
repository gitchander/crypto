package kalyna

const (
	bitsPerByte    = 8
	bytesPerUint64 = 8

	bytesPerWord = bytesPerUint64
)

// Block words size.
const (
	wordsPerBlock128 = 2
	wordsPerBlock256 = 4
	wordsPerBlock512 = 8
)

// Key words size.
const (
	wordsPerKey128 = 2
	wordsPerKey256 = 4
	wordsPerKey512 = 8
)

// Number of enciphering rounds size depending on key length.
const (
	roundsForKey128 = 10
	roundsForKey256 = 14
	roundsForKey512 = 18
)

// Block bytes sizes.
const (
	bytesPerBlock128 = wordsPerBlock128 * bytesPerWord
	bytesPerBlock256 = wordsPerBlock256 * bytesPerWord
	bytesPerBlock512 = wordsPerBlock512 * bytesPerWord
)

// Key bytes size.
const (
	bytesPerKey128 = wordsPerKey128 * bytesPerWord
	bytesPerKey256 = wordsPerKey256 * bytesPerWord
	bytesPerKey512 = wordsPerKey512 * bytesPerWord
)

const (
	reductionPolynomial = 0x011d // x^8 + x^4 + x^3 + x^2 + 1

	reductionPolynomialUint8 = uint8(reductionPolynomial & 0xff)
)

const (
	bytesPoolSizeMax = bytesPerBlock512
	wordsPoolSizeMax = bytesPerBlock512
)

const (
	BlockSize128 = 128 / bitsPerByte
	BlockSize256 = 256 / bitsPerByte
	BlockSize512 = 512 / bitsPerByte

	KeySize128 = 128 / bitsPerByte
	KeySize256 = 256 / bitsPerByte
	KeySize512 = 512 / bitsPerByte
)
