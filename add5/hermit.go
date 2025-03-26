package main

type Hermite struct {
	zs    []float64
	coefs []float64
}

func NewHermite(xs []float64, ys [][]float64) *Hermite {
	zs := make([]float64, 0)
	indeces := make([]int, 0)

	for i, y := range ys {
		for range y {
			zs = append(zs, xs[i])
			indeces = append(indeces, i)
		}
	}

	dynamic := make([][]float64, len(zs))
	for i := range len(dynamic) {
		if i != 0 {
			dynamic[i] = make([]float64, len(dynamic)-i)
		}
	}

	for i, y := range ys {
		for range y {
			dynamic[0] = append(dynamic[0], ys[i][0])
		}
	}

	for l := range len(dynamic) - 1 {
		for i := 0; i+l+1 < len(dynamic); i++ {
			if indeces[i] == indeces[i+l+1] {
				dynamic[l+1][i] = ys[indeces[i]][l+1] / float64(factorial(l+1))
			} else {
				dynamic[l+1][i] = (dynamic[l][i+1] - dynamic[l][i]) / (xs[indeces[i+l+1]] - xs[indeces[i]])
			}
		}
	}

	coefs := make([]float64, 0)
	for l := range dynamic {
		coefs = append(coefs, dynamic[l][0])
	}

	return &Hermite{
		zs:    zs,
		coefs: coefs,
	}
}

func factorial(x int) int {
	if x <= 0 {
		return 1
	}

	return x * factorial(x-1)
}

func (h *Hermite) Evaluate(x float64) float64 {
	mult := 1.0
	res := 0.0
	for i, coef := range h.coefs {
		res += coef * mult
		mult *= (x - h.zs[i])
	}
	return res
}
