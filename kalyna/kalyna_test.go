package kalyna

import (
	"errors"
	"testing"
)

type sample struct {
	id         string
	config     Config
	key        []uint64
	plaintext  []uint64
	ciphertext []uint64
}

var samples = []sample{
	{
		id: "B.2.6-encryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize128,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
		},
		plaintext: []uint64{
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
		},
		ciphertext: []uint64{
			0x20ac9b777d1cbf81, 0x06add2b439eac9e1,
		},
	},
	{
		id: "B.2.6-decryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize128,
		},
		key: []uint64{
			0x08090a0b0c0d0e0f, 0x0001020304050607,
		},
		plaintext: []uint64{
			0x84c70c472bef9172, 0xd7da733930c2096f,
		},
		ciphertext: []uint64{
			0x18191a1b1c1d1e1f, 0x1011121314151617,
		},
	},
	{
		id: "B.2.7-encryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize256,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
		},
		plaintext: []uint64{
			0x2726252423222120, 0x2f2e2d2c2b2a2928,
		},
		ciphertext: []uint64{
			0x8a150010093eec58, 0x144f336f16f74811,
		},
	},
	{
		id: "B.2.7-decryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize256,
		},
		key: []uint64{
			0x18191a1b1c1d1e1f, 0x1011121314151617,
			0x08090a0b0c0d0e0f, 0x0001020304050607,
		},
		plaintext: []uint64{
			0xe1dffdce56b46df3, 0x96d9ca30705f5bb4,
		},
		ciphertext: []uint64{
			0x28292a2b2c2d2e2f, 0x2021222324252627,
		},
	},
	{
		id: "B.2.8-encryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize256,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
		},
		plaintext: []uint64{
			0x2726252423222120, 0x2f2e2d2c2b2a2928,
			0x3736353433323130, 0x3f3e3d3c3b3a3938,
		},
		ciphertext: []uint64{
			0x3521c90e573d6ef6, 0x8c2abddc23e3daae,
			0x5a0d6a20ec6339a0, 0x2cd97f61245c3888,
		},
	},
	{
		id: "B.2.8-decryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize256,
		},
		key: []uint64{
			0x18191a1b1c1d1e1f, 0x1011121314151617,
			0x08090a0b0c0d0e0f, 0x0001020304050607,
		},
		plaintext: []uint64{
			0x864e67967823c57f, 0xa34b8b3fb0e9c103,
			0xd3c33f2c597c5bab, 0xe30fb28625d1ed61,
		},
		ciphertext: []uint64{
			0x38393a3b3c3d3e3f, 0x3031323334353637,
			0x28292a2b2c2d2e2f, 0x2021222324252627,
		},
	},
	{
		id: "B.2.9-encryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize512,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
			0x2726252423222120, 0x2f2e2d2c2b2a2928,
			0x3736353433323130, 0x3f3e3d3c3b3a3938,
		},
		plaintext: []uint64{
			0x4746454443424140, 0x4f4e4d4c4b4a4948,
			0x5756555453525150, 0x5f5e5d5c5b5a5958,
		},
		ciphertext: []uint64{
			0x7ab6b7e6e9906960, 0xb76822d793d8d64b,
			0x02e1d73c3cc8028e, 0xd95dfefda8742efd,
		},
	},
	{
		id: "B.2.9-decryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize512,
		},
		key: []uint64{
			0x38393a3b3c3d3e3f, 0x3031323334353637,
			0x28292a2b2c2d2e2f, 0x2021222324252627,
			0x18191a1b1c1d1e1f, 0x1011121314151617,
			0x08090a0b0c0d0e0f, 0x0001020304050607,
		},
		plaintext: []uint64{
			0x82d4da67277a3118, 0x078d78a1b907cdbc,
			0x97845f9e1898705e, 0xe06aba796d910b2d,
		},
		ciphertext: []uint64{
			0x58595a5b5c5d5e5f, 0x5051525354555657,
			0x48494a4b4c4d4e4f, 0x4041424344454647,
		},
	},
	{
		id: "B.2.10-encryption",
		config: Config{
			BlockSize: BlockSize512,
			KeySize:   KeySize512,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
			0x2726252423222120, 0x2f2e2d2c2b2a2928,
			0x3736353433323130, 0x3f3e3d3c3b3a3938,
		},
		plaintext: []uint64{
			0x4746454443424140, 0x4f4e4d4c4b4a4948,
			0x5756555453525150, 0x5f5e5d5c5b5a5958,
			0x6766656463626160, 0x6f6e6d6c6b6a6968,
			0x7776757473727170, 0x7f7e7d7c7b7a7978,
		},
		ciphertext: []uint64{
			0x6a351c811be3264a, 0x1a239605cad61da6,
			0xa1f347aa5483ba67, 0xb856eb20c3ee1d3e,
			0x66ab5b1717f4d095, 0x6cc815bb34f1d62f,
			0xb7fe6e85266a90cb, 0xd9d90d947264bcc5,
		},
	},
	{
		id: "B.2.10-decryption",
		config: Config{
			BlockSize: BlockSize512,
			KeySize:   KeySize512,
		},
		key: []uint64{
			0x38393a3b3c3d3e3f, 0x3031323334353637,
			0x28292a2b2c2d2e2f, 0x2021222324252627,
			0x18191a1b1c1d1e1f, 0x1011121314151617,
			0x08090a0b0c0d0e0f, 0x0001020304050607,
		},
		plaintext: []uint64{
			0x5252a025338480ce, 0x29d8a9e614d7ea1b,
			0xbd45a8e90e1e38fd, 0xa346fad954450492,
			0xf2b13b85dbef7f75, 0x6ae6753b839dff97,
			0xdc1b29b5ab5741af, 0x22ff5aaa13bb94f0,
		},
		ciphertext: []uint64{
			0x78797a7b7c7d7e7f, 0x7071727374757677,
			0x68696a6b6c6d6e6f, 0x6061626364656667,
			0x58595a5b5c5d5e5f, 0x5051525354555657,
			0x48494a4b4c4d4e4f, 0x4041424344454647,
		},
	},
}

func TestSamples(t *testing.T) {
	for _, v := range samples {
		err := testSample(t, v)
		if err != nil {
			t.Fatalf("sample %q error: %s", v.id, err)
		} else {
			t.Logf("sample %q success", v.id)
		}
	}
}

func testSample(t *testing.T, v sample) error {

	k, err := newKalynaContext(v.config.BlockSize, v.config.KeySize)
	if err != nil {
		return err
	}

	err = k.KeyExpand(v.key)
	if err != nil {
		return err
	}

	ciphertext := make([]uint64, len(v.ciphertext))
	k.Encipher(v.plaintext, ciphertext)

	if !compareWords(ciphertext, v.ciphertext) {
		return errors.New("failed enciphering")
	}
	//t.Log("Success enciphering")

	plaintext := make([]uint64, len(v.plaintext))
	k.Decipher(ciphertext, plaintext)

	if !compareWords(plaintext, v.plaintext) {
		return errors.New("failed deciphering")
	}
	//t.Log("Success deciphering")

	//t.Log("plaintext:", wordsToString(plaintext))
	//t.Log("ciphertext:", wordsToString(ciphertext))

	return nil
}

func compareWords(a, b []uint64) bool {
	n := len(a)
	if len(b) != n {
		return false
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
