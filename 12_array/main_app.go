// array dalam bahasa pemrograman Go
// setiap array hanya bisa menampung satu tipe data
package main

import "fmt"

func main() {
	// deklarasi array
	// var = [jumlah data]tipe data{nilai1, nilai2, nilai3}
	// bisa jumlah value lebih sedikit dari jumlah data
	// tapi jumlah data tidak bisa lebih sedikit dari jumlah value
	var witch= [3]string{"Alice", "Nicole", "Rhinnedottir"}

	// menampilkan array
	fmt.Println(witch)
	// menampilkan berdasarkan index
	fmt.Println(witch[0])
	fmt.Println(witch[1])
	fmt.Println(witch[2])

	// menampilkan panjang array. bukan jumlah value
	// tapi jumlah data yang bisa ditampung array
	fmt.Println(len(witch))

	// mengubah value array
	witch[1] = "Reeyn"
	fmt.Println(witch)
}