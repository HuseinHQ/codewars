package dnatorna

// import "fmt"

// func DNAtoRNA(dna string) string {
// 	var result string;
// 	for _, char := range dna {
// 		if char == 'T' {
// 			result = fmt.Sprintf("%s%c", result, 'U')
// 			} else {
// 				result = fmt.Sprintf("%s%c", result, char)
// 		}
// 	}
// 	return result
// }

// Other formula
import "strings"

// func DNAtoRNA(dna string) string {
//   return strings.Replace(dna, "T", "U", -1) // -1 maksudnya jumlah character yang direplace, klo <0 berarti unlimited, klo 2 berarti ya 2 character aja
// }

func DNAtoRNA(dna string) string {
  return strings.Replace(dna, "T", "U", -1)
}