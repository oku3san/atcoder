package main

import (
	"fmt"
	"math"
)

func main() {
	var x, a, b float64
	fmt.Scan(&x, &a, &b)

	if math.Abs(a-x) > math.Abs(b-x) {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
