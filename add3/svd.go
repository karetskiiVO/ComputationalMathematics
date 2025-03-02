package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func SVD(A *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense) {
	r, c := A.Dims()
	
	if r > c {
		U, Sigma, V := SVD(mat.DenseCopyOf(A.T()))
		return V, mat.DenseCopyOf(Sigma.T()), mat.DenseCopyOf(U.T())
	}

	U := mat.NewDense(r, r, nil)
	Sigma := mat.NewDense(r, c, nil)
	V := mat.NewDense(c, c, nil)
	
	// вычисляем Sigma
	AAT := &mat.Dense{}
	AAT.Mul(A, A.T())

	lambdas, _ := SymmetricPowerMethod(AAT, 1000, 1e-8)
	for i := range min(r, c) {
		Sigma.Set(i, i, math.Sqrt(lambdas[i]))
	}

	// вычисляем V 
	ATA := &mat.Dense{}
	ATA.Mul(A.T(), A)

	vs := make([]*mat.VecDense, 0, c)

	for _, lambda := range lambdas {
		B := mat.DenseCopyOf(ATA) // B = A * A^T - lambda * I
		for i := range c {
			B.Set(i, i, B.At(i, i) - lambda)
		}

		lu := &mat.LU{}
		lu.Factorize(B)
		U := &mat.TriDense{}
		lu.UTo(U)

		vs = append(vs, mat.NewVecDense(c, nil))

		x := vs[len(vs)-1].RawVector().Data
		for i := c - 1; i >= 0; i-- {
			if math.Abs(U.At(i, i)) < 1e-8 {
				x[i] = 1
			} else {
				sum := 0.0
				for j := i + 1; j < c; j++ {
					sum += U.At(i, j) * x[j]
				}
				x[i] = -sum / U.At(i, i)
			}
		}

		v := mat.NewVecDense(c, x)
		v.ScaleVec(1/mat.Norm(v, 2), v)
	} 
	
	ToBasis(&vs, c, 1e-8)
	for i := range c {
		V.SetCol(i, vs[i].RawVector().Data)
	}

	// вычисляем U
	us := make([]*mat.VecDense, r)

	for i := range r {
		us[i] = &mat.VecDense{}
		us[i].MulVec(A, vs[i])
		us[i].ScaleVec(1 / math.Sqrt(lambdas[i]), us[i])
	}

	for i := range r {
		U.SetCol(i, us[i].RawVector().Data)
	}

	return U, Sigma, V
}

func ToBasis (vects *[](*mat.VecDense), dim int, eps float64) {
	for i := range dim {
		if len(*vects) == dim {
			return
		}
		
		v := mat.NewVecDense(dim, nil)
		v.SetVec(i, 1)

		for _, vect := range *vects {
			proj := mat.Dot(v, vect)
			v.AddScaledVec(v, -proj, vect)
			
			mag := v.Norm(2)
			if mag < eps {
				break
			}

			v.ScaleVec(1/mag, v)
		}

		mag := v.Norm(2)
		if mag > eps {
			(*vects) = append((*vects), v)
		}
	}
}