package main

import (
	"github.com/shogo82148/float16"
)

func AbsoluteSum(numbers ...float16.Float16) float64 {
	var res float64

	for _, val := range numbers {
		res += val.Float64()
	}

	return res
}

func SimpleSum(numbers ...float16.Float16) float64 {
	var res float16.Float16

	for _, val := range numbers {
		res = res.Add(val)
	}

	return res.Float64()
}

func TreeSum(numbers ...float16.Float16) float64 {
	if len(numbers) == 0 {
		return 0
	}

	var treeSum func(numbers ...float16.Float16) float16.Float16
	treeSum = func(numbers ...float16.Float16) float16.Float16 {
		if len(numbers) == 1 {
			return numbers[0]
		}

		halfLen := (len(numbers) + 1) / 2

		return treeSum(numbers[:halfLen]...).Add(treeSum(numbers[halfLen:]...))
	}

	return treeSum(numbers...).Float64()
}

func KahanSum(numbers ...float16.Float16) float64 {
	var res, rate float16.Float16

	for _, val := range numbers {
		y := val.Sub(rate)
		tmp := res.Add(y)

		rate = (tmp.Sub(res)).Sub(y)
		res = tmp
	}

	return res.Float64()
}
