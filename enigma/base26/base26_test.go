package base26

import (
	"bytes"
	"testing"

	"github.com/gitchander/crypto/utils/random"
)

func TestRandom(t *testing.T) {

	r := random.NewRandNow()
	const n = 100
	data := make([]byte, n)

	for i := 0; i < 10000; i++ {

		m := r.Intn(n + 1)
		as := data[:m]

		random.FillBytes(r, as)

		es := make([]byte, EncodedLenMax(len(as)))
		n := Encode(es, as)
		es = es[:n]

		bs := make([]byte, DecodedLenMax(len(es)))
		n, err := Decode(bs, es)
		if err != nil {
			t.Fatal(err)
		}
		bs = bs[:n]

		if !(bytes.Equal(as, bs)) {
			t.Logf("%s: [%x]", "as", as)
			t.Logf("%s: [%x]", "bs", bs)
			t.Fatal("samples is not equal")
		}
	}
}

func TestStringRandom(t *testing.T) {

	r := random.NewRandNow()
	const n = 100
	data := make([]byte, n)

	for i := 0; i < 10000; i++ {

		m := r.Intn(n + 1)
		as := data[:m]

		random.FillBytes(r, as)

		es := EncodeToString(as)
		bs, err := DecodeString(es)
		if err != nil {
			t.Fatal(err)
		}

		if !(bytes.Equal(as, bs)) {
			t.Logf("%s: [%x]", "as", as)
			t.Logf("%s: [%x]", "bs", bs)
			t.Fatal("samples is not equal")
		}
	}
}

// todo
// func TestStringABC(t *testing.T) {

// 	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 	r := random.NewRandNow()
// 	const n = 10
// 	data := make([]byte, n)

// 	for i := 0; i < 1000; i++ {
// 		m := r.Intn(n + 1)
// 		es := data[:m]
// 		for i := range es {
// 			es[i] = alphabet[r.Intn(len(alphabet))]
// 		}

// 		//t.Logf("%s", es)

// 		as := make([]byte, DecodedLenMax(len(es)))
// 		n, err := Decode(as, es)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		as = as[:n]

// 		//t.Logf("%x", as)

// 		bs := make([]byte, EncodedLenMax(len(as)))
// 		n = Encode(bs, as)
// 		bs = bs[:n]

// 		if !(bytes.Equal(es, bs)) {
// 			t.Logf("%s: %s", "es", es)
// 			t.Logf("%s: [%x]", "es", es)
// 			t.Logf("%s: [%x]", "bs", bs)
// 			t.Fatal("samples is not equal")
// 		}
// 	}
// }
