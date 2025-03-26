package main

import (
	"slices"

	"bitbucket.org/pcas/tools/hash"
	"golang.org/x/mobile/event/key"
)

type HermitPolynom struct {
}

func (poly *HermitPolynom) Fit(xs []float64, ys [][]float64) {

	n := len(xs)
	if len(ys) != n {
		panic("")
	}

	factorial := make([]uint64, n)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = uint64(i) * factorial[i-1]
	}

	ktotal := 0
	ks := make([]int, n)
	for i := range n {
		ks[i] = len(ys[i])
		ktotal += ks[i]
	}

	dynamic := make([]*UniversalMap[[]int, float64], ktotal)
	for i := range dynamic {
		dynamic[i] = NewUniversalMap[[]int, float64](
			slices.Equal[[]int],
			func(key []int) uintptr {
				return uintptr(hash.IntSlice(key))
			},
		)
	}

	for i, k := range ks {
		for j_ := range k {
			j := j_ + 1
			key := make([]int, n)
			key[i] = j

			dynamic[j].Insert(key, ys[i][j] / float64(factorial[j]))
		}
	}

}

func (poly *HermitPolynom) Predict(x float64) float64 {
	return 0
}
