package main

import "fmt"

func main() {
	nama := "Agunng"

	if nama == "Agung" {
		fmt.Println("Selamat datang", nama)
	} else if nama == "Budi" {
		fmt.Println("Selamat datang", nama)
	} else {
		fmt.Println("Boleh kenalan", nama, "?")
	}

	//short statment
	if length := len(nama); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
