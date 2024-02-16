package reversedsequence

func ReverseSeq(n int) []int {
	var result []int
	for i := n; i > 0; i-- {
		result = append(result, i)
	}

  return result
}

// other solution
// func ReverseSeq(n int) []int {
//   arr := make([]int, n)
//   for i:= 0; i < n; i++ {
//     arr[i] = n - i
//   }
//   return arr
// }