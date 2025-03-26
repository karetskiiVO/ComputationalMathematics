package main

import (
	"math"

	"github.com/vorduin/slices"
	"golang.org/x/exp/constraints"

	"gonum.org/v1/gonum/mat"
)

type Remez struct {
	coefs []float64
}

func MakeRemez(xmin, xmax float64, degree int, f func(float64) float64, eps float64, limit int) *Remez {
	subseg := 5
	if limit == -1 {
		limit = math.MaxInt
	}

	xs := make([]float64, degree+2)
	for i := range xs {
		xs[i] = xmin + (xmax-xmin)*float64(i)/float64(degree+1)
	}

	n := len(xs)

	A := mat.NewDense(n, n, nil)

	var coefs []float64
	for range limit {
		for i := range n {
			val := 1.0
			for j := n - 2; j >= 0; j-- {
				A.Set(i, j, val)
				val *= xs[i]
			}

			sign := 1.0
			if i%2 == 1 {
				sign = -1.0
			}
			A.Set(i, n-1, sign)
		}
		fs := mat.NewVecDense(n, slices.Map(f, xs))
		sol := &mat.VecDense{}
		sol.SolveVec(A, fs)
		coefs = slices.Reverse(sol.RawVector().Data[0 : sol.Len()-1])

		diffFunc := func(x float64) float64 {
			mul := 1.0
			res := 0.0

			for _, c := range coefs {
				res += c * mul
				mul *= x
			}

			return math.Abs(res - f(x))
		}

		maxx := 0.0
		diffmax := 0.0
		for i := range len(xs) - 1 {
			for range subseg {
				x := xs[i] + (xs[i+1]-xs[i])*float64(i)/float64(subseg-1)

				diff := diffFunc(x)
				if diff >= diffmax {
					diffmax = diff
					maxx = x
				}
			}
		}

		idx := FindClosestIndexSortedGeneric(xs, maxx)
		xs[idx] = maxx
	}

	return &Remez{
		coefs: coefs,
	}
}

func (r *Remez) Evaluate(x float64) float64 {
	mul := 1.0
	res := 0.0

	for _, c := range r.coefs {
		res += c * mul
		mul *= x
	}

	return res
}

func FindClosestIndexSortedGeneric[T constraints.Integer | constraints.Float](slice []T, target T) int {
	n := len(slice)
	if n == 0 {
		return -1
	}

	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low > 0 {
		diffCurrent := diff(slice[low], target)
		diffLeft := diff(slice[low-1], target)
		if diffLeft < diffCurrent {
			return low - 1
		}
	}
	return low
}

func diff[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}
