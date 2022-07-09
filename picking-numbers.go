package main

import "fmt"

//https://www.hackerrank.com/challenges/picking-numbers/problem?isFullScreen=false

func main() {
	a := []int32{4, 6, 5, 3, 3, 1}
	fmt.Println(pickingNumbers(a))
}
func pickingNumbers(a []int32) int32 {
	/*her bir sayı için seçilen  sayıdan en büyük  1 büyük ya da 1 küçük olan sayıların adedini döndürür.*/
	var final []int32
	for i := 0; i < len(a); i++ {
		var arrayup []int32
		var arraydown []int32
		number := a[i]
		for m := 0; m < len(a); m++ {
			if a[m] == number || a[m] == number+1 {
				arrayup = append(arrayup, a[m])
			}
			if a[m] == number || a[m] == number-1 {
				arraydown = append(arraydown, a[m])
			}
			if len(arrayup) > len(final) {
				final = arrayup
			} else if len(arraydown) > len(final) {
				final = arraydown
			}
		}
		fmt.Println(arrayup, "-----", arraydown)
	}
	return int32(len(final))
}
