package reducebutgrow

func Grow(arr []int) int{
	var result int

	for i, val := range arr {
		if i == 0 {
			result = 1
		}
		result = result * val
	}

	return result
}

// other version
// func Grow(arr []int) int{
//   result := 1
  
//   for _, n := range arr {
//     result *= n
//   }
  
//   return result
// }