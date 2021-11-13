package main

import (
	"fmt"
)

func modifyVariable(val *string) {
	*val = "You"
}

func main() {
	var target string = "World"
	var indirect_target *string = &target

	fmt.Printf("value=%v\n", target)
	fmt.Printf("ptrs=%v\n", indirect_target)
	fmt.Printf("ptrs_value=%v\n", *indirect_target)

	modifyVariable(&target)

	fmt.Printf("value=%v\n", target)
	fmt.Printf("ptrs=%v\n", indirect_target)
	fmt.Printf("ptrs_value=%v\n", *indirect_target)
}

