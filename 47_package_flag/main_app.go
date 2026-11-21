// package flag dalam bahasa pemrograman Go
// digunakan untuk mem-parsing (mengurai) argumen atau
// opsi baris perintah (command-line flags) saat program dijalankan dari terminal
package main

import (
	"flag"
	"fmt"
)

func main() {
	// mendaftarkan spesifikasi opsi ke memori dan alokasi pointer berisi nilai default
	hostname := flag.String("hostname", "localhost", "put your database hostname")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")
	number := flag.Int("number", 0, "put your number")

	// jika go run biasa, maka yang akan di parse adalah value default
	// jika go run file.go -namaflag=value, maka value flag juga akan berubah
	flag.Parse()
	// untuk mengakses value dari flag menggunakan pointer, *namaflag
	fmt.Println("hostname:", *hostname)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("number:", *number)
}
