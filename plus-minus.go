package main

import "fmt"

//dizideki + - ve 0 ların tüm diziye oranını float olarak yazdır.

func main() {
	var array = []int{1, 1, 0, -1, -1}
	var positive, negative, zero int

	for _, j := range array {
		if j > 0 {
			positive++
		} else if j < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Println(float64(positive)/float64(len(array)), float64(negative)/float64(len(array)), float64(zero)/float64(len(array)))
}
