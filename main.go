package main

import "fmt"

func main() {
	fmt.Println("app started")
	fmt.Println("result: ", add(3, 4))
}

func add(a int, b int) int {
	return a + b
}
