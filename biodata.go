package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	NoAbsen   int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var students = []Student{
	{NoAbsen: 1, Nama: "Hamonangan Sitorus", Alamat: "Toba", Pekerjaan: "Back-End Engineer I", Alasan: "Golang is Awesome"},
	{NoAbsen: 2, Nama: "Esra Manurung", Alamat: "Jakarta", Pekerjaan: "Back-End Engineer I", Alasan: "Golang is Good"},
	{NoAbsen: 3, Nama: "Saut Sihotang", Alamat: "Bekasi", Pekerjaan: "Back-End Engineer I", Alasan: "Golang is Cool"},
	{NoAbsen: 23, Nama: "Teguh Ainul", Alamat: "Surabaya", Pekerjaan: "Back-End Engineer I", Alasan: "Golang is Fast"},
	{NoAbsen: 10, Nama: "Reyhan", Alamat: "Jakarta", Pekerjaan: "Back-End Engineer I", Alasan: "Golang is Simple"},
}

func (s Student) printData() {
	fmt.Println("Nama :", s.Nama)
	fmt.Println("Alamat :", s.Alamat)
	fmt.Println("Pekerjaan :", s.Pekerjaan)
	fmt.Println("Alasan :", s.Alasan)
}

func searchStudent(noAbsen string) {

	var temp int

	absen, _ := strconv.Atoi(noAbsen)

	for index, value := range students {
		if absen == value.NoAbsen {
			temp = index + 1
		}
	}

	if temp > 0 {
		students[temp-1].printData()
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func main() {

	input := os.Args[1]

	searchStudent(input)
}
