package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	sumsArray := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		var sum int
		selected := array[i]

		for k := 0; k < len(array); k++ {
			if selected == array[k] {
				continue
			}
			sum += array[k]
		}
		sumsArray[i] = sum
	}
	fmt.Println(sumsArray)

	i := 0
	selectedItem := sumsArray[i]
	for ; i < len(sumsArray); i++ {
		if sumsArray[i] > selectedItem {
			selectedItem = sumsArray[i]
		}
	}
	biggest := selectedItem

	i = 0
	selectedItem = sumsArray[i]
	for ; i < len(sumsArray); i++ {
		if sumsArray[i] < selectedItem {
			selectedItem = sumsArray[i]
		}
	}
	smallest := selectedItem
	fmt.Println("biggest sum -> ", biggest)
	fmt.Println("smallest sum -> ", smallest)
}
