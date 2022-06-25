package main

import (
	"fmt"
	"math"
)

/*
https://www.hackerrank.com/challenges/kangaroo/problem?isFullScreen=true
*/

func main() {
	possible := kangaroo(0, 2, 5, 3)
	if possible {
		fmt.Println("possible")
	} else if !possible {
		fmt.Println("not possible")
	}
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) bool {
	step := math.Abs(float64((x1 - x2) / (v2 - v1)))
	if x1+int32(step)*v1 == x2+int32(step)*v2 {
		return true
	} else {
		return false
	}
}
