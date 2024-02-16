package arrayplusarray

func ArrayPlusArray(arr1, arr2 []int) int {
	var (
		arr1sum int
		arr2sum int
	)

	for _, num := range arr1 {
		arr1sum += num
	}

	for _, num := range arr2 {
		arr2sum += num
	}

	return arr1sum + arr2sum
}
