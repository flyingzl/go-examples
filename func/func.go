package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}
