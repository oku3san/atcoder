package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if (A+B)%3 == 0 || A%3 == 0 || B%3 == 0 {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}
