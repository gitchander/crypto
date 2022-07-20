package kalyna

import (
	"bytes"
	"errors"
	"sync"
	"testing"

	"github.com/gitchander/crypto/utils/random"
)

var configs = []Config{
	{
		BlockSize: 128,
		KeySize:   128,
	},
	{
		BlockSize: 128,
		KeySize:   256,
	},
	{
		BlockSize: 256,
		KeySize:   256,
	},
	{
		BlockSize: 256,
		KeySize:   512,
	},
	{
		BlockSize: 512,
		KeySize:   512,
	},
}

func TestRandCiphers(t *testing.T) {
	r := random.NewRandNow()
	for i := 0; i < 10; i++ {
		var (
			seed   = r.Int63()
			config = configs[r.Intn(len(configs))]
		)
		err := testSeedConfig(seed, config)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestRandCiphersSync(t *testing.T) {
	r := random.NewRandNow()
	const n = 100
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		var (
			seed   = r.Int63()
			config = configs[r.Intn(len(configs))]
		)
		go func() {
			defer wg.Done()
			err := testSeedConfig(seed, config)
			if err != nil {
				t.Log("ERR ", err)
			}
		}()
	}
	wg.Wait()
}

func testSeedConfig(seed int64, c Config) error {

	r := random.NewRandSeed(seed)

	keySize := c.KeySize / bitsPerByte
	key := make([]byte, keySize)
	random.FillBytes(r, key)

	b, err := c.NewCipher(key)
	if err != nil {
		return err
	}

	var (
		blockSize = b.BlockSize()

		plaintext  = make([]byte, blockSize)
		ciphertext = make([]byte, blockSize)

		plaintextResult = make([]byte, blockSize)
	)

	for i := 0; i < 20; i++ {

		random.FillBytes(r, plaintext)

		b.Encrypt(ciphertext, plaintext)
		b.Decrypt(plaintextResult, ciphertext)

		if !bytes.Equal(plaintext, plaintextResult) {
			return errors.New("invalid Encrypt -> Decrypt")
		}
	}

	return nil
}
