package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 10
	p := &v
	p.Y = 20
	fmt.Println(v)
}