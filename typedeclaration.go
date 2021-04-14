package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAgung NoKTP = "11111111"
	fmt.Println(ktpAgung)
	fmt.Println(NoKTP("22222222222222"))
}
