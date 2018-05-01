package main

import "fmt"

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	fmt.Println(x)

	y := make([]int, 0, 100) //returns a slice of size 0 while allocating an underlying array with size 100
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 43)
	fmt.Println(y)
	// map
	// struct
}
