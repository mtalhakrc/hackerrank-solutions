package main

import "fmt"

/*
https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem?isFullScreen=false
*/
func main() {
	result := breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	fmt.Println(result)
}

func breakingRecords(scores []int32) []int32 {
	//no, this is patrick
	highestcount := 0
	lowestcount := 0
	highest := scores[0]
	lowest := scores[0]
	for i := 1; i < len(scores); i++ {
		if scores[i] > highest {
			highestcount++
			highest = scores[i]
		} else if scores[i] < lowest {
			lowestcount++
			lowest = scores[i]
		}
	}
	return []int32{int32(lowestcount), int32(highestcount)}
}
