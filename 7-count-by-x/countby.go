package countbyx

// func CountBy(x, n int) []int {
//   var result []int
// 	for i := x; i <= n * x; i += x {
// 		if i % x == 0 {
// 			result = append(result, i)
// 		}
// 	}
// 	return result
// }

// cara lain
func CountBy(x int, n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, i*x)
	}
	return result;
}
