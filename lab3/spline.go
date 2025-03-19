package main

type Spline interface {
	Fit (xs, ys []float64)
	Predict (x float64) float64
}

