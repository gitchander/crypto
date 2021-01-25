package magma

// import "encoding/binary"

// var byteOrder = binary.LittleEndian

// func baseStep(r replacer, n *[2]uint32, x uint32) {

// 	s := n[0] + x
// 	s = r.replace(s)
// 	s = (s << 11) | (s >> 21)
// 	s ^= n[1]

// 	n[1] = n[0]
// 	n[0] = s
// }

// func encrypt(r replacer, n *[2]uint32, xs *[8]uint32) {

// 	for j := 0; j < 3; j++ {
// 		for i := 0; i < 8; i++ {
// 			baseStep(r, n, xs[i])
// 		}
// 	}
// 	for i := 8; i > 0; i-- {
// 		baseStep(r, n, xs[i-1])
// 	}

// 	swapWords(n)
// }

// func decrypt(r replacer, n *[2]uint32, xs *[8]uint32) {

// 	for i := 0; i < 8; i++ {
// 		baseStep(r, n, xs[i])
// 	}
// 	for j := 0; j < 3; j++ {
// 		for i := 8; i > 0; i-- {
// 			baseStep(r, n, xs[i-1])
// 		}
// 	}

// 	swapWords(n)
// }

// func swapWords(n *[2]uint32) {
// 	n[0], n[1] = n[1], n[0]
// }

// func getTwoUint32(src []byte, n *[2]uint32) {
// 	n[0] = byteOrder.Uint32(src[0:4])
// 	n[1] = byteOrder.Uint32(src[4:8])
// }

// func putTwoUint32(dst []byte, n *[2]uint32) {
// 	byteOrder.PutUint32(dst[0:4], n[0])
// 	byteOrder.PutUint32(dst[4:8], n[1])
// }

// func encryptBlock(r replacer, n *[2]uint32, xs *[8]uint32, dst, src []byte) {
// 	getTwoUint32(src, n)
// 	encrypt(r, n, xs)
// 	putTwoUint32(dst, n)
// }

// func decryptBlock(r replacer, n *[2]uint32, xs *[8]uint32, dst, src []byte) {
// 	getTwoUint32(src, n)
// 	decrypt(r, n, xs)
// 	putTwoUint32(dst, n)
// }
