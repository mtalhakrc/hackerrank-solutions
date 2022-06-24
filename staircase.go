package main

import "fmt"

func main() {
	var number int
	fmt.Printf("enter a number:")
	fmt.Scanf("%d", &number)

	for i := number; i > 1; i-- {
		for k := 1; k < i; k++ {
			fmt.Printf("  # ")
		}
		fmt.Printf("\n")
	}
}
