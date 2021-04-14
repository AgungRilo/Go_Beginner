package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":   "Agung",
		"alamat": "Bekasi",
	}

	person["pekerjaan"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person["pekerjaan"])
	fmt.Println(len(person))

	buku := make(map[string]string)
	buku["judul"] = "IPA"
	buku["pengarang"] = "Erlangga"
	buku["salah"] = "ups"

	fmt.Println(buku)
	fmt.Println(len(buku))

	delete(buku, "salah")

	fmt.Println(buku)
	fmt.Println(len(buku))

}
