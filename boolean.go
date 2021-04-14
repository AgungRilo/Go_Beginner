package main

import "fmt"

func main() {
	fmt.Println("Benar ", true)
	fmt.Println("Salah ", false)

	var (
		nilai             = 90
		absensi           = 80
		lulusNilai   bool = nilai > 80
		lulusAbsensi bool = absensi > 90
		lulus        bool = lulusNilai && lulusAbsensi
	)

	// fmt.Println(lulus)

	if lulus == true {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	// fmt.Println("Anda Lulus")
	// :
	// fmt.Println("Tidak Lulus")
}
