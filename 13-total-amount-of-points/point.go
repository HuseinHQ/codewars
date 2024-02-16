package point

// import (
// 	"strconv"
// 	"strings"
// )

// func Points(games []string) int {
// 	var points int
// 	for _, game := range games {
// 		scores := strings.Split(game, ":")
// 		x, _ := strconv.Atoi(scores[0])
// 		y, _ := strconv.Atoi(scores[1])

// 		if x > y {
// 			points += 3
// 		} else if x == y {
// 			points += 1
// 		}
// 	}
// 	return points
// }

// cara lain
import (
	"strings"
)

func Points(games []string) int {
	result := 0
	for _, game := range games {
		str := strings.Split(game, ":")
		x, y := str[0], str[1]
		switch {
		case x > y:
			result += 3
		case x == y:
			result += 1
		}
	}
	return result
}
