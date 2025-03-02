package main

import "gonum.org/v1/gonum/mat"

func QRLineSolver(A *mat.Dense, b *mat.VecDense) *mat.VecDense {
	Q, R := ModifiedGramSchmidt(A)
	QT := Q.T()

	Qb := &mat.VecDense{}
	Qb.MulVec(QT, b)

	n := R.RawMatrix().Cols
	x := mat.NewVecDense(n, nil)

	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += R.At(i, j) * x.AtVec(j)
		}
		x.SetVec(i, (Qb.AtVec(i)-sum)/R.At(i, i))
	}

	return x
}