package main

import "fmt"

func main() {
	//constant nilai tetap dan tidak bisa di ubah

	const firstName string = "Agung"
	const lastName = "Rilo"
	const value = 1000

	// firstName = "budi"

	//deklarasi multiple
	const (
		firstName1 = "Budi"
		lastName2  = "yono"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	fmt.Println(firstName1 + lastName2)
}
