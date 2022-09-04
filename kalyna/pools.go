package kalyna

import (
	"sync"
)

var bytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, bytesPoolSizeMax)
	},
}

func poolGetBytes(n int) []byte {
	bs := bytesPool.Get().([]byte)
	return bs[:n]
}

func poolPutBytes(bs []byte) {
	bs = bs[:0]
	bytesPool.Put(bs)
}

var wordsPool = sync.Pool{
	New: func() interface{} {
		return make([]uint64, 0, wordsPoolSizeMax)
	},
}

func poolGetWords(n int) []uint64 {
	ws := wordsPool.Get().([]uint64)
	return ws[:n]
}

func poolPutWords(ws []uint64) {
	ws = ws[:0]
	wordsPool.Put(ws)
}
