package main

import (
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func PowerMethod(A *mat.Dense, maxIter int, eps float64) (float64, *mat.VecDense) {
	n, _ := A.Dims()
	v := mat.NewVecDense(n, nil)

	if mat.Norm(A, math.Inf(1)) < eps {
		return 0, v
	}

	for {
		for i := range n {
			v.SetVec(i, rand.Float64())
		}

		Av := &mat.Dense{}
		Av.Mul(A, v)

		if mat.Norm(Av, 2) > eps {
			break
		}
	}

	lambda := float64(1)
	for range maxIter {
		Av := mat.NewVecDense(n, nil)
		Av.MulVec(A, v)

		lambdaNew := mat.Norm(Av, 2) / mat.Norm(v, 2)
		Av.ScaleVec(1/mat.Norm(Av, 2), Av)

		if math.Abs(lambdaNew-lambda) < eps {
			lambda = lambdaNew
			break
		}

		lambda = lambdaNew
		v.CopyVec(Av)
	}

	return lambda, v
}

func SymmetricPowerMethod(A *mat.Dense, maxIter int, eps float64) ([]float64, []*mat.VecDense) {
	r, c := A.Dims()
	if r != c {
		panic("matrix must be square")
	}

	lambdas := make([]float64, r)
	vectors := make([]*mat.VecDense, r)

	B := mat.DenseCopyOf(A)

	for i := range r {
		lambdas[i], vectors[i] = PowerMethod(B, maxIter, eps)
		vectors[i].ScaleVec(1/mat.Norm(vectors[i], 2), vectors[i])

		if lambdas[i] == 0 {
			break
		}

		B.RankOne(B, -lambdas[i], vectors[i], vectors[i])
	}

	for i := range r {
		if vectors[i] != nil {
			continue
		}
		vectors[i] = mat.NewVecDense(r, nil) // TODO: нормальные вектора для lambda = 0

	}

	return lambdas, vectors
}
