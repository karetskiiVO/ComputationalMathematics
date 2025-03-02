package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	A := mat.NewDense(3, 2, []float64{
		1, 2,
		3, 4,
		5, 6,
	})
	b := mat.NewVecDense(3, []float64{7, 8, 9})

	x := QRLineSolver(A, b)

	fmt.Println(mat.Formatted(x))
}