package magma

// S-box (substitution-box)
// https://en.wikipedia.org/wiki/S-box

// SBoxMagma
type ReplaceTable [8][16]byte

var RT1 = ReplaceTable{
	{0x4, 0xA, 0x9, 0x2, 0xD, 0x8, 0x0, 0xE, 0x6, 0xB, 0x1, 0xC, 0x7, 0xF, 0x5, 0x3},
	{0xE, 0xB, 0x4, 0xC, 0x6, 0xD, 0xF, 0xA, 0x2, 0x3, 0x8, 0x1, 0x0, 0x7, 0x5, 0x9},
	{0x5, 0x8, 0x1, 0xD, 0xA, 0x3, 0x4, 0x2, 0xE, 0xF, 0xC, 0x7, 0x6, 0x0, 0x9, 0xB},
	{0x7, 0xD, 0xA, 0x1, 0x0, 0x8, 0x9, 0xF, 0xE, 0x4, 0x6, 0xC, 0xB, 0x2, 0x5, 0x3},
	{0x6, 0xC, 0x7, 0x1, 0x5, 0xF, 0xD, 0x8, 0x4, 0xA, 0x9, 0xE, 0x0, 0x3, 0xB, 0x2},
	{0x4, 0xB, 0xA, 0x0, 0x7, 0x2, 0x1, 0xD, 0x3, 0x6, 0x8, 0x5, 0x9, 0xC, 0xF, 0xE},
	{0xD, 0xB, 0x4, 0x1, 0x3, 0xF, 0x5, 0x9, 0x0, 0xA, 0xE, 0x7, 0x6, 0x8, 0x2, 0xC},
	{0x1, 0xF, 0xD, 0x0, 0x5, 0x7, 0xA, 0x4, 0x9, 0x2, 0x3, 0xE, 0x6, 0xB, 0x8, 0xC},
}

var RT2 = ReplaceTable{
	{0xC, 0x4, 0x6, 0x2, 0xA, 0x5, 0xB, 0x9, 0xE, 0x8, 0xD, 0x7, 0x0, 0x3, 0xF, 0x1},
	{0x6, 0x8, 0x2, 0x3, 0x9, 0xA, 0x5, 0xC, 0x1, 0xE, 0x4, 0x7, 0xB, 0xD, 0x0, 0xF},
	{0xB, 0x3, 0x5, 0x8, 0x2, 0xF, 0xA, 0xD, 0xE, 0x1, 0x7, 0x4, 0xC, 0x9, 0x6, 0x0},
	{0xC, 0x8, 0x2, 0x1, 0xD, 0x4, 0xF, 0x6, 0x7, 0x0, 0xA, 0x5, 0x3, 0xE, 0x9, 0xB},
	{0x7, 0xF, 0x5, 0xA, 0x8, 0x1, 0x6, 0xD, 0x0, 0x9, 0x3, 0xE, 0xB, 0x4, 0x2, 0xC},
	{0x5, 0xD, 0xF, 0x6, 0x9, 0x2, 0xC, 0xA, 0xB, 0x7, 0x8, 0x1, 0x4, 0x3, 0xE, 0x0},
	{0x8, 0xE, 0x2, 0x5, 0x6, 0x9, 0x1, 0xC, 0xF, 0x4, 0xB, 0x0, 0xD, 0xA, 0x3, 0x7},
	{0x1, 0x7, 0xE, 0xD, 0x0, 0x5, 0x8, 0x3, 0x4, 0xF, 0xA, 0x6, 0x9, 0xC, 0xB, 0x2},
}

var RT3 = ReplaceTable{
	{0x4, 0x2, 0xF, 0x5, 0x9, 0x1, 0x0, 0x8, 0xE, 0x3, 0xB, 0xC, 0xD, 0x7, 0xA, 0x6},
	{0xC, 0x9, 0xF, 0xE, 0x8, 0x1, 0x3, 0xA, 0x2, 0x7, 0x4, 0xD, 0x6, 0x0, 0xB, 0x5},
	{0xD, 0x8, 0xE, 0xC, 0x7, 0x3, 0x9, 0xA, 0x1, 0x5, 0x2, 0x4, 0x6, 0xF, 0x0, 0xB},
	{0xE, 0x9, 0xB, 0x2, 0x5, 0xF, 0x7, 0x1, 0x0, 0xD, 0xC, 0x6, 0xA, 0x4, 0x3, 0x8},
	{0x3, 0xE, 0x5, 0x9, 0x6, 0x8, 0x0, 0xD, 0xA, 0xB, 0x7, 0xC, 0x2, 0x1, 0xF, 0x4},
	{0x8, 0xF, 0x6, 0xB, 0x1, 0x9, 0xC, 0x5, 0xD, 0x3, 0x7, 0xA, 0x0, 0xE, 0x2, 0x4},
	{0x9, 0xB, 0xC, 0x0, 0x3, 0x6, 0x7, 0x5, 0x4, 0x8, 0xE, 0xF, 0x1, 0xA, 0x2, 0xD},
	{0xC, 0x6, 0x5, 0x2, 0xB, 0x0, 0x9, 0xD, 0x3, 0xE, 0x7, 0xA, 0xF, 0x4, 0x1, 0x8},
}

//------------------------------------------------------------------------------
func checkValidReplaceTable(rt ReplaceTable) error {

	var x [16]int

	for _, bs := range rt {

		// reset
		for j, _ := range x {
			x[j] = 0
		}

		for j := 0; j < 16; j++ {
			b := bs[j]
			if (b >= 0) && (b < 16) {
				x[b]++
			} else {
				return ErrInvalidReplaceTable
			}
		}

		for j, _ := range x {
			if x[j] != 1 {
				return ErrInvalidReplaceTable
			}
		}
	}

	return nil
}

//------------------------------------------------------------------------------
type replacer interface {
	replace(uint32) uint32
}

func makeReplacer(rt ReplaceTable) replacer {
	if false {
		return makeReplaceTable8x16(rt)
	} else {
		return makeReplaceTable4x256(rt)
	}
}

//------------------------------------------------------------------------------
type replaceTable8x16 [8][16]byte

func makeReplaceTable8x16(rt ReplaceTable) replacer {
	v := replaceTable8x16(rt)
	return &v
}

func (p *replaceTable8x16) replace(x uint32) uint32 {
	return replaceByTable8x16(p, x)
}

func replaceByTable8x16(rt *replaceTable8x16, s0 uint32) (s1 uint32) {
	for i := 0; i < 8; i++ {
		var (
			shift = 4 * i
			j     = ((s0 >> shift) & 0xF)
		)
		s1 |= uint32(rt[i][j]) << shift
	}
	return s1
}

//------------------------------------------------------------------------------
type replaceTable4x256 [4][256]byte

func makeReplaceTable4x256(rt ReplaceTable) replacer {

	var v replaceTable4x256

	for i := range v {
		for j := 0; j < 16; j++ {
			for k := 0; k < 16; k++ {
				v[i][(j*16)+k] = rt[i*2][k] | (rt[(i*2)+1][j] << 4)
			}
		}
	}

	return &v
}

func (p *replaceTable4x256) replace(x uint32) uint32 {
	return replaceByTable4x256(p, x)
}

func replaceByTable4x256(rt *replaceTable4x256, s0 uint32) (s1 uint32) {
	for i := 0; i < 4; i++ {
		var (
			shift = 8 * i
			j     = ((s0 >> shift) & 0xFF)
		)
		s1 |= uint32(rt[i][j]) << shift
	}
	return s1
}
