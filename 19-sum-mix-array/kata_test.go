package kata

import (
	"fmt"
	"strconv"
	"testing"
)

func SumMix(arr []any) int {
	total := 0
	for _, val := range arr {
		switch v := val.(type) {
		case int:
			total += v
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				total += i
			}
		}
	}
	return total
}

func TestSumMix(*testing.T) {
	fmt.Println(SumMix([]any{1, 2, "3"}))
}
