package kalyna

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"
)

type cipherSample struct {
	id         string
	config     Config
	key        []string
	plaintext  []string
	ciphertext []string
}

var cipherSamples = []cipherSample{
	{
		id: "B.2.6-encryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize128,
		},
		key: []string{
			"000102030405060708090A0B0C0D0E0F",
		},
		plaintext: []string{
			"101112131415161718191A1B1C1D1E1F",
		},
		ciphertext: []string{
			"81BF1C7D779BAC20E1C9EA39B4D2AD06",
		},
	},
	{
		id: "B.2.6-decryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize128,
		},
		key: []string{
			"0F0E0D0C0B0A09080706050403020100",
		},
		plaintext: []string{
			"7291EF2B470CC7846F09C2303973DAD7",
		},
		ciphertext: []string{
			"1F1E1D1C1B1A19181716151413121110",
		},
	},
	{
		id: "B.2.7-encryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize256,
		},
		key: []string{
			"000102030405060708090A0B0C0D0E0F",
			"101112131415161718191A1B1C1D1E1F",
		},
		plaintext: []string{
			"202122232425262728292A2B2C2D2E2F",
		},
		ciphertext: []string{
			"58EC3E091000158A1148F7166F334F14",
		},
	},
	{
		id: "B.2.7-decryption",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize256,
		},
		key: []string{
			"1F1E1D1C1B1A19181716151413121110",
			"0F0E0D0C0B0A09080706050403020100",
		},
		plaintext: []string{
			"F36DB456CEFDDFE1B45B5F7030CAD996",
		},
		ciphertext: []string{
			"2F2E2D2C2B2A29282726252423222120",
		},
	},
	{
		id: "B.2.8-encryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize256,
		},
		key: []string{
			"000102030405060708090A0B0C0D0E0F",
			"101112131415161718191A1B1C1D1E1F",
		},
		plaintext: []string{
			"202122232425262728292A2B2C2D2E2F",
			"303132333435363738393A3B3C3D3E3F",
		},
		ciphertext: []string{
			"F66E3D570EC92135AEDAE323DCBD2A8C",
			"A03963EC206A0D5A88385C24617FD92C",
		},
	},
	{
		id: "B.2.8-decryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize256,
		},
		key: []string{
			"1F1E1D1C1B1A19181716151413121110",
			"0F0E0D0C0B0A09080706050403020100",
		},
		plaintext: []string{
			"7FC5237896674E8603C1E9B03F8B4BA3",
			"AB5B7C592C3FC3D361EDD12586B20FE3",
		},
		ciphertext: []string{
			"3F3E3D3C3B3A39383736353433323130",
			"2F2E2D2C2B2A29282726252423222120",
		},
	},
	{
		id: "B.2.9-encryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize512,
		},
		key: []string{
			"000102030405060708090A0B0C0D0E0F",
			"101112131415161718191A1B1C1D1E1F",
			"202122232425262728292A2B2C2D2E2F",
			"303132333435363738393A3B3C3D3E3F",
		},
		plaintext: []string{
			"404142434445464748494A4B4C4D4E4F",
			"505152535455565758595A5B5C5D5E5F",
		},
		ciphertext: []string{
			"606990E9E6B7B67A4BD6D893D72268B7",
			"8E02C83C3CD7E102FD2E74A8FDFE5DD9",
		},
	},
	{
		id: "B.2.9-decryption",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize512,
		},
		key: []string{
			"3F3E3D3C3B3A39383736353433323130",
			"2F2E2D2C2B2A29282726252423222120",
			"1F1E1D1C1B1A19181716151413121110",
			"0F0E0D0C0B0A09080706050403020100",
		},
		plaintext: []string{
			"18317A2767DAD482BCCD07B9A1788D07",
			"5E7098189E5F84972D0B916D79BA6AE0",
		},
		ciphertext: []string{
			"5F5E5D5C5B5A59585756555453525150",
			"4F4E4D4C4B4A49484746454443424140",
		},
	},
	{
		id: "B.2.10-encryption",
		config: Config{
			BlockSize: BlockSize512,
			KeySize:   KeySize512,
		},
		key: []string{
			"000102030405060708090A0B0C0D0E0F",
			"101112131415161718191A1B1C1D1E1F",
			"202122232425262728292A2B2C2D2E2F",
			"303132333435363738393A3B3C3D3E3F",
		},
		plaintext: []string{
			"404142434445464748494A4B4C4D4E4F",
			"505152535455565758595A5B5C5D5E5F",
			"606162636465666768696A6B6C6D6E6F",
			"707172737475767778797A7B7C7D7E7F",
		},
		ciphertext: []string{
			"4A26E31B811C356AA61DD6CA0596231A",
			"67BA8354AA47F3A13E1DEEC320EB56B8",
			"95D0F417175BAB662FD6F134BB15C86C",
			"CB906A26856EFEB7C5BC6472940DD9D9",
		},
	},
	{
		id: "B.2.10-decryption",
		config: Config{
			BlockSize: BlockSize512,
			KeySize:   KeySize512,
		},
		key: []string{
			"3F3E3D3C3B3A39383736353433323130",
			"2F2E2D2C2B2A29282726252423222120",
			"1F1E1D1C1B1A19181716151413121110",
			"0F0E0D0C0B0A09080706050403020100",
		},
		plaintext: []string{
			"CE80843325A052521BEAD714E6A9D829",
			"FD381E0EE9A845BD92044554D9FA46A3",
			"757FEFDB853BB1F297FF9D833B75E66A",
			"AF4157ABB5291BDCF094BB13AA5AFF22",
		},
		ciphertext: []string{
			"7F7E7D7C7B7A79787776757473727170",
			"6F6E6D6C6B6A69686766656463626160",
			"5F5E5D5C5B5A59585756555453525150",
			"4F4E4D4C4B4A49484746454443424140",
		},
	},
}

func TestCipherSamples(t *testing.T) {
	for _, v := range cipherSamples {
		err := testCipherSample(t, v)
		if err != nil {
			t.Fatalf("sample %q error: %s", v.id, err)
		} else {
			t.Logf("sample %q success", v.id)
		}
	}
}

func testCipherSample(t *testing.T, v cipherSample) error {

	key, err := decodeHex(v.key)
	if err != nil {
		return err
	}

	plaintext, err := decodeHex(v.plaintext)
	if err != nil {
		return err
	}

	ciphertext, err := decodeHex(v.ciphertext)
	if err != nil {
		return err
	}

	b, err := v.config.NewCipher(key)
	if err != nil {
		return err
	}

	blockSize := v.config.BlockSize
	if b.BlockSize() != blockSize {
		return fmt.Errorf("invalid blockSize: have %d, want %d",
			b.BlockSize(), blockSize)
	}

	ciphertextResult := make([]byte, len(ciphertext))
	b.Encrypt(ciphertextResult, plaintext)

	if !bytes.Equal(ciphertextResult, ciphertext) {
		return errors.New("failed enciphering")
	}

	plaintextResult := make([]byte, len(plaintext))
	b.Decrypt(plaintextResult, ciphertextResult)

	if !bytes.Equal(plaintextResult, plaintext) {
		return errors.New("failed deciphering")
	}

	return nil
}

func decodeHex(vs []string) ([]byte, error) {
	var bs []byte
	for _, v := range vs {
		data, err := hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
		bs = append(bs, data...)
	}
	return bs, nil
}
