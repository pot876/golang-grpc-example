package fibo

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math/big"
)

func FromTo(fr, to uint64) ([]string, error) {
	if to <= fr {
		return []string{}, nil
	}

	fiboIterator := FiboIterator(big.NewInt(0), big.NewInt(1))

	var resultInsertIndex uint64 = 0
	var result = make([]string, to-fr)

	if fr < 2 {
		for i := fr; i < min(2, fr+(to-fr)); i++ {
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

func keyAsString(key uint64) string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(key))

	return base32.HexEncoding.EncodeToString(b)
}
func min(a, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}
