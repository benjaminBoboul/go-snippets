package main

import "fmt"

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Height * r.Width
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect(w=%v, h=%v)", r.Width, r.Height)
}

func main() {
	var r Rect = Rect{2, 3}
	fmt.Printf("Rect area=%v\n", r.Area())
	fmt.Println(r)
}
