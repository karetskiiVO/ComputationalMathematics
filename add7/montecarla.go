package main

import "math/rand"

func KahanSum(numbers []float64) float64 {
	var sum, err, y, t float64

	for _, num := range numbers {
		y = num - err
		t = sum + y
		err = (t - sum) - y
		sum = t
	}

	return sum
}

func MMKIntegrate(f func(float64) float64, xmin, xmax, ymin, ymax float64, iter int) float64 {
	res := 0
	r := rand.New(rand.NewSource(42))

	for range iter {
		x := r.Float64()*(xmax-xmin) + xmin
		y := r.Float64()*(ymax-ymin) + ymin

		if y < f(x) {
			res++
		}
	}

	return (float64(res)/float64(iter)*(ymax-ymin) + ymin) * (xmax - xmin)
}

func Integrate(f func(float64) float64, xmin, xmax float64, iter int) float64 {
	ys := make([]float64, iter+1)
	for i := range iter + 1 {
		ys[i] = f(xmin + (xmax-xmin)*float64(i)/float64(iter))
	}

	return (KahanSum(ys) - f(xmin) - f(xmax)) * (xmax - xmin) / float64(iter)
}
