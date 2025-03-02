package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Пример матрицы A
	A := mat.DenseCopyOf(mat.NewDense(2, 3, []float64{
		3, 2, 2,
		2, 3, -2,
	}).T())

	U, Sigma, V := SVD(A)

	fmt.Println("A:")
	fmt.Println(mat.Formatted(A))
	fmt.Println("U:")
	fmt.Println(mat.Formatted(U))
	fmt.Println("Sigma:")
	fmt.Println(mat.Formatted(Sigma))
	fmt.Println("V^T:")
	fmt.Println(mat.Formatted(V.T()))

	A.Mul(U, Sigma)
	A.Mul(A, V.T())
	fmt.Println("A:")
	fmt.Println(mat.Formatted(A))
}