package stringtoarray

import "strings"

// func StringToArray(str string) []string {
// 	return strings.Split(str, " ");
// }

// cara lain
func StringToArray(str string) []string {
	return strings.Fields(str) // Fields menghilangkan seluruh whitespace dan memotong string menjadi slice
}