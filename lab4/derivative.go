package main

import (
	"gonum.org/v1/gonum/mat"
)

func DerivativeFromSeries(x float64, fxs, fys []float64, k int) float64 {
	n := len(fxs)
	// n >= k
	xs := Map(fxs, func(x1 float64) float64 { return x1 - x })
	Acontent := make([]float64, n*n)
	for j := range n {
		buf := 1.0
		for i := range n {
			Acontent[i*n+j] = buf
			buf *= xs[j] / float64(max(i, 1))
		}
	}

	buf := 1
	x_j := 1.0
	bcontent := make([]float64, n)
	for i := range bcontent {
		if i != 0 {
			buf *= i
		}

		if i < k {
			continue
		}

		bcontent[i] = float64(buf) * x_j
		x_j *= x
		buf /= max(i-k, 1) // may be error
	}

	A := mat.NewDense(n, n, Acontent)
	b := mat.NewVecDense(n, bcontent)

	c := &mat.VecDense{}
	c.SolveVec(A, b)

	//fmt.Println(mat.Formatted(c.T()))

	return mat.Dot(c, mat.NewVecDense(n, fys))
}
