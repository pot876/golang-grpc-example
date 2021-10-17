package fibo

import (
	"fmt"
	"math/big"
)

// type FiboCache interface {
// 	// get returns: 1) index of first fibo, 2) continious fibo values with length <= n
// 	// result length < n if:
// 	// a) there are values in range but not first
// 	// b) there are values in range but not last
// 	// c) there are values in range but not in between, then only first part is returned
// 	get(anchor uint64, n int) (uint64, []string, error)
// 	set(anchor uint64, values []string) error
// 	getLeft(anchor uint64, n int) (uint64, []string, error)
// }

func FromTo(fr, to uint64) ([]string, error) {
	if to <= fr {
		return []string{}, nil
	}

	fiboIterator := FiboIterator(big.NewInt(0), big.NewInt(1))

	var resultInsertIndex uint64 = 0
	var result = make([]string, to-fr)

	if fr < 2 {
		for i := fr; i < 2; i++ {
			result[resultInsertIndex] = fmt.Sprintf("%d", i)
			resultInsertIndex++
		}
	} else {
		for i := uint64(0); i < fr-2; i++ {
			fiboIterator()
		}
	}

	for ; resultInsertIndex < (to - fr); resultInsertIndex++ {
		fiboNumber := fiboIterator()
		result[resultInsertIndex] = fiboNumber.String()
	}

	return result, nil
}

func FiboIterator(a, b *big.Int) func() *big.Int {
	return func() *big.Int {
		a, b = b, a
		return b.Add(a, b)
	}
}
