package main

import (
	"math/rand"

	"github.com/shogo82148/float16"
)

func GenerateRandomSample(n int) []float16.Float16 {
	return GenerateSample(n, rand.Int63())
}

func GenerateSample(n int, seed int64) []float16.Float16 {
	r := rand.New(rand.NewSource(int64(seed)))
	res := make([]float16.Float16, n)

	for i := range res {
		res[i] = float16.FromFloat32(r.Float32())
	}

	return res
}
