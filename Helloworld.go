package main

import (
	"fmt"
)

func main() {
	var world string = "World"
	fmt.Printf("Hello %v", world)
	fmt.Scanln(world)
	fmt.Printf("\n repeated %v", world)

}
