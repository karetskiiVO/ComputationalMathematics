package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Householder(A *mat.Dense) (*mat.Dense, *mat.Dense) {
	r, c := A.Dims()
	Q := mat.NewDense(r, c, nil)
	R := mat.NewDense(c, c, nil)
	R.Copy(A)

	for i := range min(r, c) {
		Q.Set(i, i, 1)
	}

	for j := range c {
		x := R.ColView(j).(*mat.VecDense).SliceVec(j, r).(*mat.VecDense)
		normX := mat.Norm(x, 2)

		v := mat.NewVecDense(r-j, nil)
		v.CopyVec(x)
		v.SetVec(0, v.AtVec(0)+math.Copysign(normX, x.AtVec(0)))
		normV := mat.Norm(v, 2)
		v.ScaleVec(1/normV, v)

		for k := j; k < c; k++ {
			col := R.ColView(k).(*mat.VecDense).SliceVec(j, r).(*mat.VecDense)
			dot := mat.Dot(v, col)
			col.AddScaledVec(col, -2*dot, v)
		}

		for k := range r {
			row := Q.RowView(k).(*mat.VecDense).SliceVec(j, c).(*mat.VecDense)
			dot := mat.Dot(v, row.SliceVec(0, r-j).(*mat.VecDense))
			row.AddScaledVec(row, -2*dot, v)
		}
	}

	return Q, R
}
