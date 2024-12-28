package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:=x/2
	prev:=z
	for {
		z-=(z*z-x)/(2*z)
		if math.Abs(z-prev) < 1e-5 {
			break
		}
		prev=z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
