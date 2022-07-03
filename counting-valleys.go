package main

import "fmt"

func main() {
	path := "DUDUUUUUUUUDUDDUUDUUDDDUUDDDDDUUDUUUUDDDUUUUUUUDDUDUDUUUDDDDUUDDDUDDDDUUDDUDDUUUDUUUDUUDUDUDDDDDDDDD"
	result := countingValleys(int32(len(path)), path)
	fmt.Println(result)
}

func countingValleys(steps int32, path string) int32 {
	var countvalley int32 = 0
	var altitude int32 = 0
	var list = make([]int32, steps+1)
	list[0] = 0
	var i int32 = 0
	for ; i < steps; i++ {
		if path[i] == 'U' {
			altitude++
		} else if path[i] == 'D' {
			altitude--
		}
		list[i+1] = altitude
		if altitude == 0 && list[i] < 0 {
			countvalley++
		}
	}
	return countvalley
}
