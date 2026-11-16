// struct dalam bahasa pemrograman Go
// digunakan untuk menyimpan beberapa field dan tipe datanya
// field-field itu disimpan dalam satu nilai
package main

import "fmt"

// deklarasi struct, digunakan sebagai template
type person struct {
	// bisa menyimpan berbagai jenis field dan tipe datanya
	name string
	age int
	married bool
}

func main() {
	// membuat var yang merupakan struct
	aditya := person{
		// menyimpan value dalam fieldnya
		name: "aditya",
		age: 22,
		married: false,
	}
	// menampilkan setiap field dari var struct
	fmt.Println(aditya.name)
	fmt.Println(aditya.age)
	fmt.Println(aditya.married)
	
	// mengubah value dari field
	aditya.name = "wahyu"
	fmt.Println(aditya.name)
}