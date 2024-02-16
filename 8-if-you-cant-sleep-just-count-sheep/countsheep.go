package countsheep

// import (
// 	"fmt"
// 	"strconv"
// )

// func CountSheep(num int) string {
// 	result := ""
// 	for i := 1; i <= num; i++ {
// 		result = fmt.Sprintf("%s%s sheep...",result, strconv.Itoa(i))
// 	}
// 	return result
// }

// cara lain
import (
	"fmt"
	"strings"
)

func CountSheep(num int) string {
  var sb strings.Builder

  for count := 1; count <= num; count++ {
        fmt.Fprintf(&sb, "%d sheep...", count)
  }
  
  return sb.String()
}