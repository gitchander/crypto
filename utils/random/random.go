package random

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

func FillBytes(r *rand.Rand, bs []byte) {
	const (
		bitsPerByte    = 8
		bytesPerUint64 = 8
		bitsPerUint64  = bitsPerByte * bytesPerUint64
	)
	var (
		ax uint64 // bits accumulator
		an int    // bits count
	)
	for i := range bs {
		if an < bitsPerByte {
			ax = r.Uint64()
			an = bitsPerUint64
		}
		bs[i] = byte(ax)
		ax >>= bitsPerByte
		an -= bitsPerByte
	}
}
