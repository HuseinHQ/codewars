package main

import (
	// grow "2024-coding-challenge/1-reduce-but-grow"
	// reversedsequence "2024-coding-challenge/2-reversed-sequence"
	// ishegonnasurvive "2024-coding-challenge/3-is-he-gonna-survive"
	// sentencesmash "2024-coding-challenge/4-sentence-smash"
	// willyoumakeit "2024-coding-challenge/5-will-you-make-it"
	// dnatorna "2024-coding-challenge/6-dna-to-rna"
	// countbyx "2024-coding-challenge/7-count-by-x"
	// countsheep "2024-coding-challenge/8-if-you-cant-sleep-just-count-sheep"
	// maxmin "2024-coding-challenge/9-find-maximum-and-minimum-values-of-a-list"
	// stringtoarray "2024-coding-challenge/10-convert-a-string-to-an-array"
	// rps "2024-coding-challenge/11-rock-paper-scissors"
	// arrayplusarray "2024-coding-challenge/12-array-plus-array"
	point "2024-coding-challenge/13-total-amount-of-points"
	"fmt"
)

// nomor 1
// func main() {
// 	var arr = []int{1, 2, 3, 4}
// 	result := grow.Grow(arr)

// 	fmt.Println(result)
// }

// nomor 2
// func main() {
// 	int := 5
// 	result := reversedsequence.ReverseSeq(int)

// 	fmt.Println(result)
// }

// nomor 3
// func main() {
// 	bullets, dragons := 10, 20
// 	fmt.Println(ishegonnasurvive.Hero(bullets, dragons))
// }

// nomor 4
// func main() {
// 	words := []string{"hello", "world", "this", "is", "great"}
// 	fmt.Println(sentencesmash.Smash(words))
// }

// nomor 5
// func main() {
// 	fmt.Println(willyoumakeit.ZeroFuel(50, 25, 2))
// }

// nomor 6
// func main() {
// 	fmt.Println(dnatorna.DNAtoRNA("GCAT"))
// }

// nomor 7
// func main() {
// 	fmt.Println(countbyx.CountBy(1, 10))
// 	fmt.Println(countbyx.CountBy(2, 5))
// 	fmt.Println(countbyx.CountBy(100, 5))
// }

// nomor 8
// func main() {
// 	fmt.Println(countsheep.CountSheep(3))
// }

// nomor 9
// func main() {
// 	fmt.Println("Max: ", maxmin.Max([]int{4,6,2,1,9,63,-134,566}))
// 	fmt.Println("Min: ", maxmin.Min([]int{4,6,2,1,9,63,-134,566}))
// }

// nomor 10
// func main() {
// 	fmt.Println(stringtoarray.StringToArray("I love arrays they are my favorite"))
// }

// nomor 11
// func main() {
// 	fmt.Println(rps.Rps("scissors", "paper"))
// 	fmt.Println(rps.Rps("scissors", "scissors"))
// 	fmt.Println(rps.Rps("scissors", "rock"))
// }

// nomor 12
// func main() {
// 	fmt.Println(arrayplusarray.ArrayPlusArray([]int{1, 2, 3}, []int{4, 5, 6}))
// }

// nomor 13
func main() {
	games := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	fmt.Println(point.Points(games))
}
