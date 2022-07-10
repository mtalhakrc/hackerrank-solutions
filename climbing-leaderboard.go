package main

import "fmt"

//player ve rankedleri büyükten küçüğe doğru sıra ile ilerlediğimizde, playerdaki i. elemnan i+1. elemanın levelinden daha büyük olamadığı kesindir.
//bu nedenle her bir eleman için ranked dizisini baştan sona dönmeye gerek yok.
//her baştan sona dönen döngüler için timeout hatası veriyor :(

func main() {
	ranked := []int32{100, 90, 90, 80, 75, 60}
	player := []int32{50, 65, 77, 90, 102}
	fmt.Println(climbingLeaderboard(ranked, player))
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	board := make([]int32, len(player))
	//add tmp
	ranked = append(ranked, -1)
	var level int32 = 1
	k := 0
	for i := len(player) - 1; i >= 0; i-- {
		for k < len(ranked) {
			if player[i] < ranked[k] {
				if ranked[k] == ranked[k+1] {
					k++
					continue
				}
				level++
				k++
			} else {
				break
			}
		}
		board[i] = level
	}
	return board
}

/*
BU KOD DA ÇALIŞIYOR FAKAT YAVAŞ OLDUĞU İÇİN TİMEOUT HATASI VERİYOR.
	board := make([]int32, len(player))
	var level int32
	for i := 0; i < len(player); i++ {
		level = 1
		if player[i] < ranked[0] {
			level++
		}
		for m := 1; m < len(ranked); m++ {
			if player[i] < ranked[m] {
				if ranked[m] == ranked[m-1] {
					continue
				}
				level++
			}
		}
		board[i] = level
	}
	return board





	TİMEOUT HATASI.
	final := make([]int32, len(player))
	board := make(map[int32]int32, len(ranked))
	board[ranked[0]] = 1
	for i := 1; i < len(ranked); i++ {
		if board[ranked[i]] == board[ranked[i-1]] {
			board[ranked[i]]++
		} else {
			board[ranked[i]] = board[ranked[i]] + 1
		}
	}
	var level int32
	for i := 0; i < len(player); i++ {
		level = 1
		for k, _ := range board {
			if player[i] < k {
				level++
			}
		}
		final[i] = level
	}
	return final
*/
