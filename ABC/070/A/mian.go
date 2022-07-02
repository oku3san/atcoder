package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	n := strconv.Itoa(N)

	if n[0] == n[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
