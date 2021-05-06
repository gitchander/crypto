package magma

import (
	"testing"

	"github.com/gitchander/crypto/utils/random"
)

func TestAddMod32(t *testing.T) {

	r := random.NewRandNow()

	for i := 0; i < 1000000; i++ {

		a := r.Uint32()
		b := r.Uint32()

		err := addMod32Test(a, b)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func addMod32Test(a, b uint32) error {

	var (
		s1 = add_mod32_v1(a, b)
		s2 = add_mod32_v2(a, b)
		s3 = add_mod32_v3(a, b)
	)
	const format = "wrong (%d + %d) mod (2^32)"

	if s1 != s2 {
		return newErrorf(format, a, b)
	}

	if s1 != s3 {
		return newErrorf(format, a, b)
	}

	return nil
}

func TestAddMod32M1(t *testing.T) {

	r := random.NewRandNow()

	for i := 0; i < 1000000; i++ {

		a := r.Uint32()
		b := r.Uint32()

		err := addMod32M1Test(a, b)
		if err != nil {
			t.Error(err)
		}
	}

	var rangeTest = func(min, max int64) {

		for ia := min; ia <= max; ia++ {
			for ib := min; ib <= max; ib++ {

				a := uint32(ia)
				b := uint32(ib)

				err := addMod32M1Test(a, b)
				if err != nil {
					t.Error(err)
				}
			}
		}
	}

	rangeTest(0, 1000)
	rangeTest(maxUint32-1000, maxUint32)
}

func addMod32M1Test(a, b uint32) error {
	var (
		s1 = add_mod32m1(a, b)
		s2 = add_mod32m1_sample(a, b)
	)
	if s1 != s2 {
		return newErrorf("wrong (%d + %d) mod (2^32 - 1)*", a, b)
	}
	return nil
}
