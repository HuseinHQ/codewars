package kata

import (
	"fmt"
	"sort"
	"testing"
)

func SumArray(arr []int) int {
	sort.IntSlice.Sort(arr)
	sum := 0
	for i, val := range arr {
		if i != 0 && i != len(arr)-1 {
			sum += val
		}
	}
	return sum
}

func TestSumArray(*testing.T) {
	test := SumArray([]int{1, 2, 3, 4})
	fmt.Println(test)
}
