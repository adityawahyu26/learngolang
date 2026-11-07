// operasi aritmatika dalam bahasa pemrograman Go
package main

import "fmt"

func main() {
	// deklarasi variabel
	x := 2266
	y := 112

	fmt.Println("Nilai x:", x)
	fmt.Println("Nilai y:", y)

	// operasi tambah
	tambah := x + y
	fmt.Println("Hasil Penjumlahan:", tambah)

	// operasi kurang
	kurang := x - y
	fmt.Println("Hasil Pengurangan:", kurang)

	// operasi kali
	kali := x * y
	fmt.Println("Hasil Perkalian:", kali)

	// operasi bagi
	bagi := x / y
	fmt.Println("Hasil Pembagian:", bagi)

	// operasi sisa bagi
	sisa_bagi := x % y
	fmt.Println("Hasil Sisa Bagi:", sisa_bagi)
}