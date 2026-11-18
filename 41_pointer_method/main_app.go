// pointer dalam method atau struct method di Go
// digunakan untuk mengubah atau mengakses dari method
package main

import "fmt"

// deklarasi struct
type Man struct {
	name string
}

// struct method sebagai pointer untuk mengakses data
func (man *Man) Married() {
	// perubahan disini akan memengaruhi data aslinya
	man.name = "Mr. " + man.name
} 

func main() {
	// membuat variable struct
	aditya := Man{"Aditya Wahyu"}
	// memanggil method struct
	aditya.Married()
	// datanya akan ikut berubah sesuai struct method
	fmt.Println(aditya.name)
}