// package os dalam bahasa pemrograman Go
// digunakan untuk berinteraksi dengan sistem operasi
package main

import (
	"fmt"
	"os"
)

func main() {
	// menampilkan args yang dimasukkan di dalam slice
	args := os.Args
	// ini akan menampilkan path program
	fmt.Println(args)
	// misal di pwsh, go run main_app.go aditya wahyu
	// index[0] = path program
	// fmt.Println(args[1]) -> index[1] = "aditya"
	// fmt.Println(args[2]) -> index[2] = "wahyu"

	// menampilkan hostname
	name, err := os.Hostname()
	// jika tidak error, akan menampilkan hostname
	if err == nil {
		fmt.Println("Hostname:", name)
		// jika error, akan menampilkan error
	} else {
		fmt.Println("error:", err.Error())
	}

	// mengakses environment variable
	// untuk buat env di pwsh, $env:namaenv="value"
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	// ini untuk mencetak env
	fmt.Println("username:", username)
	fmt.Println("password:", password)
}
