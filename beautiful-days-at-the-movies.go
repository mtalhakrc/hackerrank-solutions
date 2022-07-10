package main

import (
	"fmt"
	"strconv"
)

//https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem?isFullScreen=true

func main() {
	result := beautifulDays(20, 23, 6)
	fmt.Println(result)
}

func beautifulDays(i int32, j int32, k int32) int32 {
	nums := make([]int32, j-i+1)
	m := 0
	for i <= j {
		nums[m] = i
		i++
		m++
	}
	//make revert
	revertedstr := make([]string, len(nums))
	for m = 0; m < len(nums); m++ {
		str := strconv.Itoa(int(nums[m]))
		var tmp string
		for a := len(str) - 1; a >= 0; a-- {
			tmp = tmp + string(str[a])
		}
		revertedstr[m] = tmp
	}

	revertednums := make([]int32, len(revertedstr))
	for m = 0; m < len(revertedstr); m++ {
		a, _ := strconv.Atoi(revertedstr[m])
		revertednums[m] = int32(a)
	}

	//calculate difference between i-reverted(i)
	results := make([]int32, len(nums))
	for m = 0; m < len(nums); m++ {
		results[m] = abs(nums[m], revertednums[m])
	}
	var count int32 = 0
	for m = 0; m < len(results); m++ {
		if results[m]%k == 0 {
			count++
		}
	}
	return count
}

func abs(num1 int32, num2 int32) int32 {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}
