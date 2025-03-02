package main

import "gonum.org/v1/gonum/mat"

func GramSchmid(A *mat.Dense) (*mat.Dense, *mat.Dense) {
	r, c := A.Dims()

	Q := mat.NewDense(r, c, nil)
	R := mat.NewDense(c, c, nil)

	for j := range c {
		v := A.ColView(j)
		u := mat.NewVecDense(r, nil)
		u.CopyVec(v)

		for i := range j {
			q := Q.ColView(i)
			
			rVal := mat.Dot(v, q)
			R.Set(i, j, rVal)
			u.AddScaledVec(u, -rVal, q)
		}

		norm := mat.Norm(u, 2)
		R.Set(j, j, norm)
		for i := range r {
			Q.Set(i, j, u.At(i, 0)/norm)
		}
	}

	return Q, R
}

func ModifiedGramSchmidt(A *mat.Dense) (*mat.Dense, *mat.Dense) {
    r, c := A.Dims()
	
    Q := mat.NewDense(r, c, nil)
    R := mat.NewDense(c, c, nil)
    Q.Copy(A)

    for j := range c {
        R.Set(j, j, mat.Norm(Q.ColView(j), 2))

        for i := range r {
            Q.Set(i, j, Q.At(i, j)/R.At(j, j))
        }

        for k := j + 1; k < c; k++ {
            R.Set(j, k, mat.Dot(Q.ColView(j), Q.ColView(k)))

            for i := range r {
                Q.Set(i, k, Q.At(i, k)-Q.At(i, j)*R.At(j, k))
            }
        }
    }

    return Q, R
}