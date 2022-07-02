package main

import "fmt"

/*
https://www.hackerrank.com/challenges/grading/problem?isFullScreen=true

önündeki 5 katı sayısı ile grade farkı 3 ya da daha fazla ise yuvarlanmayacak
yuvarlanacak sayının yuvarlanmış hali  40 dan küçük olacaksa yuvarlanmayacak
40 geçer not
*/
const pivot int = 3

func main() {
	grades := []int{4, 73, 39, 38, 33}
	fmt.Println(grades)

	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			continue
		}
		fark := 5 - (grades[i] % 5)
		if fark < pivot {
			grades[i] = grades[i] + fark
		}
	}
	fmt.Println(grades)
}
