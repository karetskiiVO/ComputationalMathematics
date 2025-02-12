package main

import (
	"github.com/shogo82148/float16"
)

type ExperimentResult struct {
	n        int
	absolute float64
	simple   float64
	tree     float64
	kahan    float64
}

func Experiment(n int, generator func(n int) []float16.Float16) ExperimentResult {
	arr := generator(n)

	return ExperimentResult{
		n,
		AbsoluteSum(arr...),
		SimpleSum(arr...),
		TreeSum(arr...),
		KahanSum(arr...),
	}
}
