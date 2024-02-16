package maxmin

func Max(list []int) int {
	max := list[0]
	for _, num := range list {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(list []int) int {
	min := list[0]
	for _, num := range list {
		if num < min {
			min = num
		}
	}
	return min
}