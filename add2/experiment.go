package main

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type ExperimentResult struct {
	n                           int
	GramSchmidtDuration         time.Duration
	ModifiedGramSchmidtDuration time.Duration
	Householder                 time.Duration
}

func Experiment(n int) (res ExperimentResult) {
	eps := float64(1e-8)

	generator := func(n int) []float64 {
		res := make([]float64, n*n)

		for i := range res {
			res[i] = (rand.Float64() - 0.5) * 1000
		}

		return res
	}

	res.n = n
	A := mat.NewDense(n, n, generator(n))
	B := mat.NewDense(n, n, nil)

	start := time.Now()
	Q, R := GramSchmid(A)
	B.Mul(Q, R)
	B.Sub(B, A)
	if mat.Norm(B, math.Inf(+1)) > eps {
		panic("wrong Gram-Schmid")
	}
	res.GramSchmidtDuration = time.Since(start)

	start = time.Now()
	Q, R = ModifiedGramSchmidt(A)
	B.Mul(Q, R)
	B.Sub(B, A)
	if mat.Norm(B, math.Inf(+1)) > eps {
		panic("wrong Modified Gram-Schmid")
	}
	res.ModifiedGramSchmidtDuration = time.Since(start)

	start = time.Now()
	Q, R = Householder(A)
	B.Mul(Q, R)
	B.Sub(B, A)
	if mat.Norm(B, math.Inf(+1)) > eps {
		panic("Householder")
	}
	res.Householder = time.Since(start)

	return
}
