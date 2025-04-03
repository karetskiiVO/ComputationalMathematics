package main

import (
	"math"
	"math/rand"
)

// target
func Experiment(scale float64) float64 {
	size := 6
	r := rand.New(rand.NewSource(42))
	xs := Map(make([]struct{}, size), func(struct{}) float64 { return scale * (r.Float64() - 0.5) })
	ys := Map(xs, math.Exp)

	return math.Abs(DerivativeFromSeries(0, xs, ys, 2) - math.Exp(0))
}
