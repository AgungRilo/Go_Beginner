package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// bulan[5] = "Ubah"
	// fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(slice1)
}
