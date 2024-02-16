package kata

import (
	"fmt"
	"testing"
)

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3

	if avg >= 90 {
		return 'A'
	} else if avg >= 80 {
		return 'B'
	} else if avg >= 70 {
		return 'C'
	} else if avg >= 60 {
		return 'D'
	} else {
		return 'F'
	}
}

func TestGetGrade(*testing.T) {
	fmt.Println(GetGrade(90, 91, 92))
}
