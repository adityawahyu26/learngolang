// error interface dalam bahasa pemrograman Go
// untuk menampilkan error yang ada dalam program
package main

import (
	"fmt"
	"errors" // memanggil package errors
)

// func pembagian, return nya 2 yaitu type data int dan error
func divide(x int, y int) (int, error) {
	// jika pembaginya bernilai 0
	if y == 0 {
		// return int nya 0
		// return error panggil func New() dari package errors
		return 0, errors.New("tidak bisa dibagi 0")
	// jika pembaginya bukan 0, baru run operasinya
	} else {
		result := x / y
		// mengembalikan hasil operasi, return error diisi nil
		return result, nil
	}
}
func main() {
	// deklarasi var untuk menyimpan return dari func
	result, err := divide(100,20)
	// result, err := divide(100,0)

	// jika return error bernilai nil, tampilkan return result
	if err == nil {
		fmt.Println("hasil pembagian:", result)
	// jika return error ada isinya
	} else {
		// maka akan menampilkan errornya
		// memanggil func Error() dari interface error
		// untuk menampilkan pesan errornya
		fmt.Println("Error:", err.Error())
	}
}