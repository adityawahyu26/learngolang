// closure dalam bahasa pemrograman Go
// bisa memengaruhi value lain diluar function
package main

import "fmt"

func main() {
	// var sebelum di increment
	x := 11
	fmt.Println("nilai x:", x)
	// func untuk increment var
	number := func() {
		fmt.Println("value di increment +1")
		x++
		// tidak bisa jika var belum di deklarasikan
		// y++
	}
	// val y jika di increment harus dibawahnya
	// y := 17
	// fmt.Println("nilai y:", y)
	// memanggil  func yang increment var
	number()
	number()
	// menampilkan hasil akhir var setelah increment func
	fmt.Println("nilai x:", x)
}