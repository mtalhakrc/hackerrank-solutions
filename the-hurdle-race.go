package main

import "fmt"

func main() {
	height := []int32{1, 6, 3, 5, 2}
	var k int32 = 4
	fmt.Println(hurdleRace(k, height))
}
func hurdleRace(k int32, height []int32) int32 {
	var dose int32 = 0
	for i := 0; i < len(height); i++ {
		if k < height[i] {
			dose = dose + height[i] - k
			k = height[i]
		}
	}
	return dose
}
