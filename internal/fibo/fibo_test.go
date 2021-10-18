package fibo

import (
	"math/big"
	"testing"
)

func TestFromTo(t *testing.T) {
	fromTo(t, 0, 1, []string{"0"})
	fromTo(t, 0, 2, []string{"0", "1"})
	fromTo(t, 0, 3, []string{"0", "1", "1"})

	fromTo(t, 1, 2, []string{"1"})
	fromTo(t, 1, 3, []string{"1", "1"})

	fromTo(t, 2, 3, []string{"1"})
	fromTo(t, 2, 4, []string{"1", "2"})

	fromTo(t, 3, 4, []string{"2"})
	fromTo(t, 3, 5, []string{"2", "3"})

	fromTo(t, 0, 5, []string{"0", "1", "1", "2", "3"})
	fromTo(t, 1, 5, []string{"1", "1", "2", "3"})
	fromTo(t, 2, 5, []string{"1", "2", "3"})
	fromTo(t, 3, 5, []string{"2", "3"})
	fromTo(t, 4, 5, []string{"3"})
	fromTo(t, 5, 5, []string{})
	fromTo(t, 6, 5, []string{})
}

func fromTo(t *testing.T, fr uint64, to uint64, expected []string) {
	result, _ := FromTo(fr, to)
	e := eq(result, expected)

	if !e {

		t.Logf("range: %d-%d, got: %v, expected:%v", fr, to, result, expected)
		t.FailNow()
	}
}

func TestFiboIterator(t *testing.T) {
	f := FiboIterator(big.NewInt(0), big.NewInt(1))
	expectedNumbers := []int64{
		1,
		2,
		3,
		5,
		8,
		13,
		21,
		34,
		55,
		89}

	slot := big.NewInt(0)
	for _, v := range expectedNumbers {
		nextNumber := f()
		if nextNumber.Cmp(slot.SetInt64(v)) != 0 {
			t.Logf("expected %v, got %v", v, nextNumber)
			t.FailNow()
		}
	}
}

func TestKeyAsString(t *testing.T) {
	prev := ""
	for i := 0; i < 10000; i++ {
		k := keyAsString(uint64(i))
		if !(k > prev) {
			println("i: %d, expected %v > %v", i, k, prev)
			t.FailNow()
		}
		prev = k
	}
}
func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
