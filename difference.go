package main

import "fmt"

func main() {
	var fName = "Agung"
	var lName = "Rilo"

	var result bool = fName == lName
	fmt.Println("nama", result)

	var value1 = 100
	var value2 = 200

	fmt.Println("b", value1 < value2)
	fmt.Println("s", value1 > value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 >= value2)
	fmt.Println("s", value1 == value2)
	fmt.Println(value1 != value2)
}
