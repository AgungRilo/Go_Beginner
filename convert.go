package main

import "fmt"

func main() {
	//konversi int
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//konversi string
	var (
		name           = "Agung"
		a       byte   = name[0]
		eString string = string(a)
	)
	fmt.Println(name)
	fmt.Println(eString)

}
