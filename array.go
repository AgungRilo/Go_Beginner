package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Agung"
	names[2] = "Budi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [3]int{
		90,
		96,
		100,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	var lagi [10]int

	fmt.Println(len(lagi))
	//10 karena panjang data yang dihitung bukan jumlah

}
