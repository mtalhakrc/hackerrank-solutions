package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*https://www.hackerrank.com/challenges/time-conversion/problem?isFullScreen=true */

func main() {
	time := "07:05:45PM"
	asd, err := timeConversion(time)
	if err != nil {
		panic(err)
	}
	fmt.Println(asd)
}

func timeConversion(time string) (string, error) {
	hour, err := strconv.Atoi(time[:2])
	if err != nil {
		return "", err
	}
	meridian := time[8:]
	if meridian == "PM" && hour != 12 {
		newtime := strconv.Itoa(12+hour) + time[2:8]
		return newtime, err
	} else if meridian == "AM" {
		return time[:8], err
	} else {
		return "", errors.New("invalid format")
	}
}
