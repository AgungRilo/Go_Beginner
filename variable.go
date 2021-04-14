package main

import "fmt"

func main() {
	var name string

	name = "Agung Rilo"
	fmt.Println(name)

	name = "Budioyono"
	fmt.Println(name)

	var friendName = "Afif"
	fmt.Println(friendName)

	var age int8 = 30
	fmt.Println(age)

	var ages int = 30
	fmt.Println(ages)

	//Tidak Menggunakan Var agar singkat

	city := "Bekasi"
	fmt.Println(city)

	//cukup di awal menggunakan :=

	city = "Jakarta"
	fmt.Println(city)

	//deklarasi variable gabungan

	var (
		firstName = "Agung"
		lastName  = "Rilo"
	)

	fmt.Println(firstName + " " + lastName)
}
