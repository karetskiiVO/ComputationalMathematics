package main

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type ExperimentResult struct {
	n        int
	residual float64
	t        time.Duration
}

func Experiment(A *mat.Dense, b *mat.VecDense, n int, solver func(*mat.Dense, *mat.VecDense, int, float64) *mat.VecDense) ExperimentResult {
	start := time.Now()
	x := solver(A, b, n, 0)
	t := time.Since(start)

	r := &mat.VecDense{}
	r.MulVec(A, x)
	r.SubVec(b, r)

	return ExperimentResult{
		residual: mat.Norm(r, math.Inf(1)),
		t:        t,
		n:        n,
	}
}

func CreateTask(n int) (*mat.Dense, *mat.VecDense) {
	data := make([]float64, n*n)
	for i := range data {
		data[i] = (rand.Float64() - 0.5) * 1000
	}

	A := mat.NewDense(n, n, data)
	A.Mul(A, A.T())
	
	vect := mat.NewVecDense(n, nil)

	for i := range n {
		vect.SetVec(i, (rand.Float64() - 0.5) * 1000)
	}

	return A, vect
}
