package main

import "fmt"

/*https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true*/

func main() {
	arr := []int32{1, 3, 2, 6, 1, 2}
	result := divisibleSumPairs(0, 3, arr)
	fmt.Println(result)
}

// n girdi sayısı ------ k toplam sayısı
func divisibleSumPairs(n int32, k int32, arr []int32) int32 {
	var count int32
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[j]+arr[i])%k == 0 {
				count++
			}
		}
	}
	return count
}
