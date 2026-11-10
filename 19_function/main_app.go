// function dalam bahasa pemrograman Go
// digunakan untuk membuat fungsi yang bisa dipakai berulang
// func(), untuk memanggilnya menggunakan func()
package main

import "fmt"

// func untuk membuat beberapa perintah
func celestia() {
	fmt.Println("hello master")
	fmt.Println("welcome to the world")
	fmt.Println("The celestia welcome you")
}

// func utama yang secara otomatis dijalankan program
func main() {
	// memanggil func tertentu 
	celestia()
}