package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum= 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}