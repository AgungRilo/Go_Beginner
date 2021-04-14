package main

import "fmt"

func main() {
	var (
		a = 3
		b = 5
	)

	a += b
	b = a - b
	a -= b

	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(b)
}
