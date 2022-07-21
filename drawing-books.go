package main

import "fmt"

/*https://www.hackerrank.com/challenges/drawing-book/problem?isFullScreen=true&h_r=next-challenge&h_v=zen*/

func main() {
	result := pageCount(6, 2)
	fmt.Println(result)
}

//int n: the number of pages in the book
//int p: the page number to turn to
func pageCount(n int32, p int32) int32 {
	//0 means empty page
	//kitap ya baştan ya da sondan zaten açık olduğu kabul edildi.
	//calculate will open page number.
	var number int32 = 2 + ((n / 2) * 2)
	book := make([]int32, number)
	//default front of book page
	book[0] = 0
	book[1] = 1

	//setting up page numbers.
	var i int32 = 2
	for ; i <= n; i++ {
		book[i] = book[i-1] + 1
	}

	var countfront int32 = 0
	i = 0
	for i <= n {
		if book[i] == p || book[i+1] == p {
			break
		} else {
			i += 2
			countfront++
		}
	}

	var countback int32 = 0
	i = number - 1
	for i >= 0 {
		if book[i] == p || book[i-1] == p {
			break
		} else {
			i -= 2
			countback++
		}
	}

	if countfront < countback {
		return countfront
	} else {
		return countback
	}
	/*
		//print
		for k, j := range book {
			fmt.Println(k, "---->", j)
		}
		fmt.Println("front--->", countfront)
		fmt.Println("back--->", countback)
		return 0
	*/
}
