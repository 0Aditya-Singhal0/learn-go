package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	// Append works on a nil slice.
	s = append(s, 42)
	printSlice(s)

	s = append(s, 13, 14, 15)
	printSlice(s)
}
