package main

import "fmt"

func main() {
	fmt.Println(saveThePrisoner(4, 6, 4))
}

//There are n  prisoners, 2 pieces of candy and distribution  starts at s chair

func saveThePrisoner(n int32, m int32, s int32) int32 {
	var result int32
	var final int32
	if m > n {
		result = m % n
		if result == 0 {
			result = n
		}
		final = result + s - 1
		final = final % n
		if final == 0 {
			final = n
		}
	} else if m <= n {
		final = m + s - 1
		final = final % n
		if final == 0 {
			final = n
		}
	}
	return final
}
