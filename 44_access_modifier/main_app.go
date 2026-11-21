// access modifier dalam bahasa pemrograman Go
// jika ingin func bisa diakses dari package lain
// maka huruf depan namanya harus kapital
package main

import (
	// import package
	"learngolang/44_access_modifier/Hello"
	"fmt"
)

func main() {
	// func ini berhasil dipanggil dari package lain
	morning := Hello.GoodMorning("aditya wahyu")
	fmt.Println(morning)

	// func dari package lain ini akan gagal dipanggil
	// night := Hello.goodNight("aditya wahyu")
	// fmt.Println(night)
}