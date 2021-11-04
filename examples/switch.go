package main

import (
  "fmt"
)

func main() {
  weekDay := 2
  
  switch weekDay {
    case 1:
      fmt.Println("Start learning Golang.")
    case 2:
      fmt.Println("Able to proper use structures.")
    case 3,4,5:
      fmt.Println("Realize it was Rust from the beginning.")
    default:
      fmt.Println("Finally, inner peace.")
  }
}
