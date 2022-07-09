package main

import "fmt"

//https://www.hackerrank.com/challenges/cats-and-a-mouse/problem?isFullScreen=true&h_r=next-challenge&h_v=zen
func main() {
	result := catAndMouse(1, 3, 2)
	fmt.Println(result)
}
func catAndMouse(x int32, y int32, z int32) string {
	for {
		if x == z && y == z {
			return "Mouse C"
		}
		if x < z {
			x++
		} else if x > z {
			x--
		} else {
			return "Cat A"
		}
		if y < z {
			y++
		} else if y > z {
			y--
		} else {
			return "Cat B"
		}
	}
}
