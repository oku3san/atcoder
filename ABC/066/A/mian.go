package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	arr := make([]int, 3)

	for i := range arr {
		fmt.Scan(&num)
		arr[i] = num
	}

	sort.Ints(arr)
	fmt.Println(arr[0] + arr[1])

}
