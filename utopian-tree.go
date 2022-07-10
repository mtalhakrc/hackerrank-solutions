package main

import "fmt"

func main() {
	fmt.Println(utopianTree(5))
}
func utopianTree(n int32) int32 {
	var height int32 = 1
	var tmp int32 = 0
	for tmp <= n {
		if tmp == 0 {
			tmp++
			continue
		}
		if tmp%2 == 1 {
			height *= 2
		} else if tmp%2 == 0 {
			height += 1
		}
		tmp++
	}
	return height
}
