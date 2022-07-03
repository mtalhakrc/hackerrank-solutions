package main

import "fmt"

/*https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true*/
func main() {
	path := "DDUUDDUDUUUD"
	result := countingValleys(int32(len(path)), path)
	fmt.Println(result)
	//expected -> 1

}

//U -> up   D -> down
/*
	//pivota göre sağdaki pivot  kendisinden önceki sayı kendisinden küçükse, vadi geçmiş olur. büyük ise dağ geçmiş olur.
	//iki pivot arasında 1 adımdan fazla yol gidilmişse yeni pivot belirlenir.
*/
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
	}
	//fmt.Println("orjinal list ---> ", list)
	countvalley = calculatevalley(list, countvalley)
	return countvalley
}

func calculatevalley(list []int32, countvalley int32) int32 {
	var i int32 = 1
	var oldpivotindex int32 = 0
	var pivotindex int32 = 0
	oldpivotindex = pivotindex
	for ; i < int32(len(list)); i++ {
		if list[i] == list[pivotindex] {
			pivotindex = i
			if list[pivotindex-1] < list[pivotindex] {
				countvalley++
			}
			//yeni pivot?
			if (i+oldpivotindex)/2 > 1 {
				var newlist = make([]int32, len(list[pivotindex:pivotindex]))
				_ = copy(newlist, list[oldpivotindex+1:pivotindex])
				fmt.Println(newlist)
				countvalley = calculatevalley(newlist, countvalley)
			}
		}
	}
	return countvalley
}
