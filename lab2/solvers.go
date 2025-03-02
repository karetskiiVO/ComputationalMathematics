package main

import "gonum.org/v1/gonum/mat"

func SteepestSolver(A *mat.Dense, b *mat.VecDense, maxIter int, tol float64) *mat.VecDense {
	n := b.Len()
	x := mat.NewVecDense(n, nil)

	r := mat.NewVecDense(n, nil)
	Ar := mat.NewVecDense(n, nil)

	for range maxIter {
		r.MulVec(A, x)
		r.SubVec(b, r)

		residualNorm := mat.Norm(r, 2)
		if residualNorm < tol {
			break
		}

		Ar.MulVec(A, r)
		tau := mat.Dot(r, r) / mat.Dot(r, Ar)

		x.AddScaledVec(x, tau, r)
	}

	return x
}

func MinimalSolver(A *mat.Dense, b *mat.VecDense, maxIter int, tol float64) *mat.VecDense {
	n := b.Len()
	x := mat.NewVecDense(n, nil)

	r := mat.NewVecDense(n, nil)
	Ar := mat.NewVecDense(n, nil)

	for range maxIter {
		r.MulVec(A, x)
		r.SubVec(b, r)

		residualNorm := mat.Norm(r, 2)
		if residualNorm < tol {
			break
		}

		Ar.MulVec(A, r)
		tau := mat.Dot(r, Ar) / mat.Dot(Ar, Ar)

		x.AddScaledVec(x, tau, r)
	}

	return x
}
