package main

import "fmt"

/*https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true*/

func main() {
	socks := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(int32(len(socks)), socks))
}

//int n: the number of socks in the pile
//int ar[n]: the colors of each sock
func sockMerchant(n int32, ar []int32) int32 {
	socks := make(map[int32]int32, n)
	var i int32 = 0
	for ; i < n; i++ {
		socks[ar[i]]++
	}
	var count int32 = 0
	for _, j := range socks {
		count += j / 2
	}
	return count
}
