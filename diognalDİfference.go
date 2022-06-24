package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("enter %d.lines %d.element:", i, j)
			fmt.Scanf("%d", &matrix[i][j])
			if matrix[i][j] > 100 {
				panic("input cannot be greater than 100")
			}
		}
	}
	fmt.Println(dioganlDifference(matrix))

}

func dioganlDifference(matrix [][]int) float64 {
	var firstSum, secondSum int

	for i, k := 0, 2; i < 3; i++ {
		firstSum += matrix[i][i]
		secondSum += matrix[i][k]
		k--

	}
	fmt.Println(firstSum, secondSum)
	return math.Abs(float64(firstSum - secondSum))
}
