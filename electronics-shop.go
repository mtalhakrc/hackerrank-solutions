package main

import "fmt"

/*https://www.hackerrank.com/challenges/electronics-shop/problem?isFullScreen=true*/
func main() {
	drives := []int32{4}
	keyboards := []int32{5}
	result := getMoneySpent(keyboards, drives, 5)
	fmt.Println(result)
}

//int keyboards[n]: the keyboard prices
//int drives[m]: the drive prices
//int b: the budget
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var list []int32
	var i, k int32 = 0, 0
	for ; i < int32(len(drives)); i++ {
		for ; k < int32(len(keyboards)); k++ {
			if drives[i]+keyboards[k] > b {
				continue
			}
			//list[i] = drives[i] + keyboards[k]
			list = append(list, drives[i]+keyboards[k])

		}
		k = 0
	}
	if list == nil {
		return -1
	}
	closest := b - list[0]
	var final int32
	for i = 1; i < int32(len(list)); i++ {
		if b-list[i] < closest {
			closest = b - list[i]
			final = list[i]
		}
	}
	return final
}
