package main

//https://www.hackerrank.com/challenges/strange-advertising/problem?isFullScreen=true

func main() {
	viralAdvertising(5)
}
func viralAdvertising(n int32) int32 {
	var cumulative int32
	var liked int32
	var shared int32
	var tmp int32 = 1
	shared = 5
	for tmp <= n {
		liked = shared / 2
		cumulative = cumulative + liked
		shared = liked * 3
		tmp++
	}
	return cumulative
}
