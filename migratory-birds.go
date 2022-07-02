package main

import "fmt"

/*https://www.hackerrank.com/challenges/migratory-birds/problem?isFullScreen=true*/

func main() {
	arr := []int32{1, 4, 4, 4, 5, 3, 3, 3, 2, 2, 2, 2, 2}
	xd := migratoryBirds(arr)
	fmt.Println(xd)
}

func migratoryBirds(arr []int32) int32 {
	birds := make(map[int32]int32, len(arr))
	for i := 0; i < len(arr); i++ {
		birds[arr[i]]++
	}

	var biggest = birds[0]
	for _, k := range birds {
		if k > biggest {
			biggest = k
		}
	}

	i := 0
	var results []int32
	for k, j := range birds {
		if j == biggest {
			results = append(results, k)
			i++
		}
	}
	i = 0
	var final = results[0]
	for i = 0; i < len(results); i++ {
		if results[i] < final {
			final = results[i]
		}
	}
	return final
}
