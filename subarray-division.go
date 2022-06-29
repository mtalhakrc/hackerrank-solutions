package main

import "fmt"

/*
https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=false
*/

func main() {
	s := []int32{1, 2, 1, 3, 2}
	var d int32 = 3
	var m int32 = 2
	result := birthday(s, d, m)
	fmt.Println(result)
}

func birthday(s []int32, d int32, m int32) int32 {
	var sum int
	count := 0
	i := 0
	for ; i < len(s)-int(m)+1; i++ {
		for k, q := i, i; q < k+int(m); q++ {
			sum = sum + int(s[q])
		}
		if sum == int(d) {
			count++
		}
		sum = 0
	}
	return int32(count)
}
