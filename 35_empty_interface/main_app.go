// empty interface / interface{} / any dalam Go
// menyimpan type data jenis apapun
package main

import "fmt"

// case sederhana, idealnya pakai switch 
// func yang menerima type data parameter apapun
func cekInput(x interface{}) {
	// if dan elif statement untuk conn dari parameter
	if x == "aditya" {
		fmt.Println("tipe data", x, "adalah string")
	} else if x == 26 {
		fmt.Println("tipe data", x, "adalah integer")
	} else if x == true || x == false {
		fmt.Println("tipe data", x, "adalah boolean")
	} else {
		fmt.Println("tipe data", x, "adalah tipe data lainnya")
	}
}

/** menggunakan switch, lebih ideal untuk menentukan type data
func cekInput(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("tipe data", v, "adalah string")
	case int:
		fmt.Println("tipe data", v, "adalah integer")
	case bool:
		fmt.Println("tipe data", v, "adalah boolean")
	default:
		fmt.Println("tipe data", v, "adalah tipe data lainnya")
	}
} 
*/

func main() {
	// memanggil func dan mengirim params bertype apapun
	cekInput("aditya")
	cekInput(26)
	cekInput(false)
}