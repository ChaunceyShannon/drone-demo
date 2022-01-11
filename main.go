package main

import (
	"fmt"
)

func plus(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("1 + 2 =", plus(1, 2))
}
