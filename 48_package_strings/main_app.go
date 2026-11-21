// package strings dalam bahasa pemrograman Go
// digunakan untuk mengakses dan memanipulasi string
package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains digunakan untuk mengecek apa ada string yang dicari
	fmt.Println(strings.Contains("aditya wahyu", "aditya"))

	// memisahkan string berdasarkan separatornya, hasilnya berupa slice
	fmt.Println(strings.Split("aditya wahyu", " "))

	// membuat seluruh string jadi lowercase
	fmt.Println(strings.ToLower("Aditya Wahyu"))

	// membuat seluruh string jadi uppercase
	fmt.Println(strings.ToTitle("aditya wahyu"))

	// jika tidak ada string dikanan dan kirinya akan potong sesuai cutsetnya
	// jika ada string lainnya, akan dilewati dan mencari ujung string
	fmt.Println(strings.Trim("     aditya wahyu xy   ", " "))

	// mengganti semua string lama yang diinginkan ke string baru
	fmt.Println(strings.ReplaceAll("aditya xy aditya", "aditya", "wahyu"))
}
