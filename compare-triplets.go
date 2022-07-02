package main

import "fmt"

func main() {
	a := make([]int, 3)
	b := make([]int, 3)

	i := 0
	for ; i < 3; i++ {
		fmt.Printf("%dth alice point:", i)
		fmt.Scanf("%d", &a[i])

		fmt.Printf("%dth bob point:", i)
		fmt.Scanf("%d", &b[i])
	}
	resp := calcScores(a, b)
	fmt.Println(resp)
}

func calcScores(a []int, b []int) []int {
	var scoreAlice, scoreBob int

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			scoreAlice++
		} else if a[i] < b[i] {
			scoreBob++
		} else {
			scoreAlice++
			scoreBob++
		}
	}
	return []int{scoreAlice, scoreBob}
}
