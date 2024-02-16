package kata

import (
	"fmt"
	"testing"
)

func GetAverage(marks []int) int {
	total := 0
	for _, value := range marks {
		total += value
	}
	avg := total / len(marks)
	return avg
}

func TestGetAverage(*testing.T) {
	fmt.Println(GetAverage([]int{2, 2, 3, 2, 2}))
}
