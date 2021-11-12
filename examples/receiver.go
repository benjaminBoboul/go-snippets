package main

import "fmt"

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Height * r.Width
}

func main() {
	var r Rect = Rect{2, 3}
	fmt.Printf("Rect area=%v\n", r.Area())
	fmt.Println(r)
}
