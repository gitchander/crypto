package crygo

import (
	"bytes"
	"testing"
)

func TestBlockCipher(t *testing.T) {

	r := newRand()

	key := make([]byte, 32)

	s1 := make([]byte, BlockSize)
	s2 := make([]byte, BlockSize)
	s3 := make([]byte, BlockSize)

	table := NewTableDefault()

	for i := 0; i < 100; i++ {

		r.FillBytes(key)

		c, err := NewBlockCipher(table, key)
		if err != nil {
			t.Error(err)
			return
		}

		for j := 0; j < 1000; j++ {

			r.FillBytes(s1)

			c.Encrypt(s2, s1)
			if bytes.Compare(s1, s2) == 0 {
				t.Error("Encrypt compare true")
				return
			}

			c.Decrypt(s3, s2)
			if bytes.Compare(s1, s3) != 0 {
				t.Error("Decrypt compare false")
				return
			}
		}
	}
}

func TestSamples(t *testing.T) {

	type Sample struct {
		Input  []byte
		Output []byte
	}

	type TestSamples struct {
		Table   []byte
		Key     []byte
		Samples []Sample
	}

	var ts = TestSamples{
		Table: table2,
		Key: []byte{
			0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88,
			0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
			0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8,
			0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xd0,
		},
		Samples: []Sample{

			Sample{
				Input:  []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
				Output: []byte{0xce, 0x5a, 0x5e, 0xd7, 0xe0, 0x57, 0x7a, 0x5f},
			},

			Sample{
				Input:  []byte{0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8},
				Output: []byte{0xd0, 0xcc, 0x85, 0xce, 0x31, 0x63, 0x5b, 0x8b},
			},
		},
	}

	table, err := NewTable(ts.Table)
	if err != nil {
		t.Error(err)
		return
	}

	c, err := NewBlockCipher(table, ts.Key)
	if err != nil {
		t.Error(err)
		return
	}

	dst := make([]byte, BlockSize)

	for _, sample := range ts.Samples {
		c.Encrypt(dst, sample.Input)
		if bytes.Compare(dst, sample.Output) != 0 {
			t.Error("not compare")
		}
	}
}

func TestTables(t *testing.T) {

	r := newRand()

	type Tables struct {
		t1 Table
		t2 Table
	}

	ts := []Tables{
		Tables{
			newReplaceTable4(table1),
			newReplaceTable8(table1),
		},
		Tables{
			newReplaceTable4(table2),
			newReplaceTable8(table2),
		},
	}

	for _, y := range ts {
		for i := 0; i < 1000000; i++ {
			u := r.Uint32()
			if y.t1.Replace(u) != y.t2.Replace(u) {
				t.Error("wrong mix")
				return
			}
		}
	}
}
