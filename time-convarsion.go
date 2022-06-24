package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*https://www.hackerrank.com/challenges/time-conversion/problem?isFullScreen=true */

/*
buna yeninden BAK
*/
func main() {
	//timePM := "12:01:00:PM"
	timeAM := "12:01:AM"
	seperated := strings.Split(timeAM, ":")
	fmt.Println(seperated)

	if seperated[2] == "AM" || seperated[2] == "PM" {
		panic("xd:)")
	}

	hour, _ := strconv.Atoi(seperated[0])
	min, _ := strconv.Atoi(seperated[1])

	if seperated[2] == "AM" {

	} else {

	}

}
