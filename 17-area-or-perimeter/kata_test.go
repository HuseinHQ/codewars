package kata

import (
	"fmt"
	"testing"
)

func AreaOrPerimeter(l, w int) int {
	if l == w {
		return l * w
	} else {
		return 2 * (l + w)
	}
}

func TestAreaOrPerimeter(*testing.T) {
	fmt.Println(AreaOrPerimeter(3, 3))
	fmt.Println(AreaOrPerimeter(6, 10))
}
