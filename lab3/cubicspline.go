package main

import (
	"slices"
)

type CubicFreeSpline struct {
	xs    []float64
	ys    []float64
	coefs []struct{ a, b, c, d float64 }
}

var _ Spline = &CubicFreeSpline{}

func (spl *CubicFreeSpline) Fit(xs, ys []float64) {
	if len(xs) != len(ys) {
		panic("mistmatch len(xs) and len(ys)")
	}

	n := len(xs)
	if n < 3 {
		panic("mistmatch len(xs) < 4")
	}

	spl.xs = slices.Clone(xs)
	spl.ys = slices.Clone(ys)
	spl.coefs = make([]struct {
		a, b, c, d float64
	}, n-1)
	dxs := make([]float64, n-1)

	b := make([]float64, n)
	alpha := make([]float64, n)
	beta_ := make([]float64, n)
	gamma := make([]float64, n)

	for i := range n - 1 {
		dxs[i] = xs[i+1] - xs[i]
	}

	for i_ := range n - 2 {
		i := i_ + 1

		b[i] = 3 * (dxs[i]*((ys[i]-ys[i-1])/dxs[i-1]) + dxs[i-1]*((ys[i+1]-ys[i])/dxs[i]))
	}
	b[0] = ((dxs[0]+2*(xs[2]-xs[0]))*dxs[1]*((ys[1]-ys[0])/dxs[0]) + dxs[0]*dxs[0]*((ys[2]-ys[1])/dxs[1])) / (xs[2] - xs[0])
	b[n-1] = (dxs[n-2]*dxs[n-2])*((ys[n-2]-ys[n-3])/dxs[n-3]) + (2*(xs[n-1]-xs[n-3])+dxs[n-2])*dxs[n-3]*((ys[n-1]-ys[n-2])/dxs[n-2])/(xs[n-1]-xs[n-3])

	beta_[0] = dxs[1]
	gamma[0] = xs[2] - xs[0]
	beta_[n-1] = dxs[n-2]
	alpha[n-1] = xs[n-1] - xs[n-3]

	for i_ := range n - 2 {
		i := i_ + 1
		beta_[i] = 2 * (dxs[i] + dxs[i-1])
		gamma[i] = dxs[i]
		alpha[i] = dxs[i-1]
	}

	c := float64(0)

	for i := range n - 1 {
		c = beta_[i]
		b[i] /= c
		beta_[i] /= c
		gamma[i] /= c

		c = alpha[i+1]
		b[i+1] -= c * b[i]
		alpha[i+1] -= c * beta_[i]
		beta_[i+1] -= c * gamma[i]
	}

	b[n-1] /= beta_[n-1]
	beta_[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		c = gamma[i]
		b[i] -= c * b[i+1]
		gamma[i] -= c * beta_[i]
	}

	for i := range n - 1 {
		c1 := (ys[i+1]-ys[i])/(dxs[i]*dxs[i]) - b[i]/dxs[i]
		c2 := b[i+1]/dxs[i] - (ys[i+1]-ys[i])/(dxs[i]*dxs[i])

		spl.coefs[i] = struct{ a, b, c, d float64 }{
			a: (c2 - c1) / dxs[i],
			b: (2.0*c1 - c2),
			c: b[i],
			d: ys[i],
		}
	}
}

func (spl *CubicFreeSpline) Predict(x float64) (y float64) {
	y = 0

	if x < spl.xs[0] {
		h := x - spl.xs[0]
		c := spl.coefs[0]
		y = c.d + h*(c.c+h*(c.b+h*c.a))
		return
	}
	if x > spl.xs[len(spl.xs)-1] {
		h := x - spl.xs[len(spl.xs)-1]
		c := spl.coefs[len(spl.coefs)-1]
		y = c.d + h*(c.c+h*(c.b+h*c.a))
		return
	}

	i, found := slices.BinarySearch(spl.xs, x)

	if found {
		y = spl.ys[i]
		return
	}

	h := x - spl.xs[i-1]
	c := spl.coefs[i-1]
	y = c.d + h*(c.c+h*(c.b+h*c.a))

	return
}
