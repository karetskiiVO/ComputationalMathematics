package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func QRLamdas(A *mat.Dense, maxIter int, eps float64) []float64 {
	isUpperTriangular := func(B *mat.Dense) bool {
		n, _ := B.Dims()

		for i := 1; i < n; i++ {
			for j := 0; j < i; j++ {
				if math.Abs(B.At(i, j)) > eps {
					return false
				}
			}
		}

		return true
	}

	n, _ := A.Dims()
	eigenvalues := make([]float64, n)

	AIter := mat.DenseCopyOf(A)

	for range maxIter {
		Q, R := ModifiedGramSchmidt(AIter)

		AIter.Mul(R, Q)
		if isUpperTriangular(AIter) {
			break
		}
	}

	for i := range n {
		eigenvalues[i] = AIter.At(i, i)
	}

	return eigenvalues
}
