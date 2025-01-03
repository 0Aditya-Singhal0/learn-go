package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x< 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n , lim float64) float64{
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(9))
	fmt.Println(sqrt(-9))	
	fmt.Println(pow(2, 10, 100))
}