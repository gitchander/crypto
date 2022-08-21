package base26

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/gitchander/crypto/utils/random"
)

func TestSamples(t *testing.T) {

	samples := []struct {
		DataHex string
		Result  string
	}{
		{
			DataHex: "A0",
			Result:  "AF",
		},
		{
			DataHex: "0123456789abcdef",
			Result:  "BYIKIHLEOKWZPO",
		},
		{
			DataHex: "50",
			Result:  "QC",
		},
		{
			DataHex: "AA",
			Result:  "KK",
		},
	}

	for _, sample := range samples {

		as, err := hex.DecodeString(sample.DataHex)
		if err != nil {
			t.Fatal(err)
		}

		//t.Logf("%x", as)

		es := make([]byte, EncodedLenMax(len(as)))
		n := Encode(es, as)
		es = es[:n]

		result := string(es)
		//t.Log("result:", result)

		if result != sample.Result {
			t.Fatalf("invalid result: have %s, want %s", result, sample.Result)
		}

		bs := make([]byte, DecodedLenMax(len(es)))
		n, err = Decode(bs, es)
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

func TestRandom(t *testing.T) {

	r := random.NewRandNow()
	const n = 100
	data := make([]byte, n)

	for i := 0; i < 1000; i++ {

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

	for i := 0; i < 1000; i++ {

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

func TestRandomOutput(t *testing.T) {

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	r := random.NewRandNow()
	const n = 100
	data := make([]byte, n)

	for i := 0; i < 1000; i++ {
		m := r.Intn(n + 1)
		es := data[:m]
		for i := range es {
			es[i] = alphabet[r.Intn(len(alphabet))]
		}

		as := make([]byte, DecodedLenMax(len(es)))
		n, err := Decode(as, es)
		if err != nil {
			// input string is invalid
			continue
		}
		as = as[:n]

		//t.Logf("%x", as)

		bs := make([]byte, EncodedLenMax(len(as)))
		n = Encode(bs, as)
		bs = bs[:n]

		if !(bytes.Equal(es, bs)) {
			t.Logf("%s: %s", "es", es)
			t.Logf("%s: [%x]", "as", as)
			t.Logf("%s: %s", "bs", bs)
			t.Fatal("samples is not equal")
		}
	}
}
