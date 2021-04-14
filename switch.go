package main

import "fmt"

func main() {
	nama := "Agung"

	switch nama {
	case "Agung":
		fmt.Println("Hello", nama)
	case "Budi":
		fmt.Println("Hello", nama)
	default:
		fmt.Println("Kenalan dong", nama)

	}
}
