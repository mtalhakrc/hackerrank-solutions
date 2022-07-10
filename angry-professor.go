package main

import "fmt"

//https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true

func main() {
	a := []int32{-1, -3, 4, 2}
	fmt.Println(angryProfessor(3, a))
}
func angryProfessor(k int32, a []int32) string {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			count++
		}
	}
	fmt.Println(count)
	if count >= int(k) {
		return "NO"
	} else {
		return "YES"
	}
}
