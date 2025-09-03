package main

import "fmt"

func main() {
	a, b := 777, 666

	fmt.Println(a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)
}
