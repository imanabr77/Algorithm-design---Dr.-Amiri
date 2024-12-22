package main

import (
	"fmt"
)

func addMatrices(matrix1, matrix2 [][]int) [][]int {
	// Check if matrices have the same dimensions
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		panic("Matrices must have the same dimensions")
	}

	// Create a new matrix to store the result
	result := make([][]int, len(matrix1))

	// Iterate through rows and columns
	for i := 0; i < len(matrix1); i++ {
		result[i] = make([]int, len(matrix1[0]))
		for j := 0; j < len(matrix1[0]); j++ {
			result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return result
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	// Define two example matrices
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrixB := [][]int{
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
	}

	// Add the matrices
	result := addMatrices(matrixA, matrixB)

	// Print the results
	fmt.Println("Matrix A:")
	printMatrix(matrixA)
	fmt.Println("\nMatrix B:")
	printMatrix(matrixB)
	fmt.Println("\nSum of Matrix A and B:")
	printMatrix(result)
}
