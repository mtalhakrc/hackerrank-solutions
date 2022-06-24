package main

import "fmt"

/*https://www.hackerrank.com/challenges/birthday-cake-candles/problem?isFullScreen=true*/

func main() {
	var candle int
	fmt.Printf("number of candles:")
	fmt.Scanf("%d", &candle)

	lenghts := make([]int, candle)
	for i := 0; i < candle; i++ {
		fmt.Printf("lenght of candle:")
		fmt.Scanf("%d", &lenghts[i])
	}

	var count int
	for i := 0; i < len(lenghts); i++ {
		biggest := FindBiggest(lenghts)
		if lenghts[i] == biggest {
			count++
		}
	}
	fmt.Println("number of longest candles:", count)
}

func FindBiggest(array []int) int {
	i := 0
	selectedItem := array[i]
	for ; i < len(array); i++ {
		if array[i] > selectedItem {
			selectedItem = array[i]
		}
	}
	return selectedItem
}

func FindSmallest(array []int) int {
	i := 0
	selectedItem := array[i]
	for ; i < len(array); i++ {
		if array[i] < selectedItem {
			selectedItem = array[i]
		}
	}
	return selectedItem
}
