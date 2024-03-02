package kata

import (
	"fmt"
	"testing"
)

func Greet(name, owner string) string {
	if name == owner {
		return "Hello boss"
	}
	return "Hello guest"
}

func TestGreet(t *testing.T) {
	fmt.Println(Greet("Husein", "Husein"))
	fmt.Println(Greet("Hasan", "Husein"))
}
