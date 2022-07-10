package main

import "fmt"

//https://www.hackerrank.com/challenges/designer-pdf-viewer/problem?isFullScreen=true

func main() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	fmt.Println(designerPdfViewer(h, "abc"))
}

func designerPdfViewer(h []int32, word string) int32 {
	var highest int32
	highest = h[word[0]-97]
	for i := 1; i < len(word); i++ {
		if h[word[i]-97] > highest {
			highest = h[word[i]-97]
		}
	}
	return int32(len(word)) * highest
}
