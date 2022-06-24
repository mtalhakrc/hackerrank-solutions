package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true

yatayda ev üzerine gelen elma ve portakal sayısını yazdır.

x ekseninde

önce ev koordinatları
sonra elma ve portakal ağaçlarının koordinatları
sonra düşecek olan elma ve portakal sayısını
sonraki satırda sırayla elma ağacından düştüğü uzaklıklar
sonraki satırda ise sırayla portakalların  portakal ağacından düştüğü uzaklıklar girilecek



*/

func main() {
	//home := make([]int, 2)
	fmt.Printf("enter home coordinates: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
	}
}
