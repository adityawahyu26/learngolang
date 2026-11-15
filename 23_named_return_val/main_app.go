// named return values pada function dalam Go
// memberikan keterangan tambahan untuk setiap return value function
package main

import "fmt"

// func namafunc() (namareturn1,.. type data) {
// jadi memudahkan mengetahui output yang dikembalikan func itu apa
func fullName() (firstName, lastName string) {
	// tidak perlu di deklarasikan, langsung tambah value
	firstName = "aditya"
	lastName = "wahyu"
	// akan otomatis return semua, jika return value diberi nama
	return
}

func main() {
	// deklarasi untuk menyimpan value
	firstName, lastName := fullName()
	// menampilkan outputnya
	fmt.Println(firstName, lastName)

	// case jika ingin menampilkan satu saja return valuenya
	// _, lastName := fullName()
	// menampilkan outputnya
	// fmt.Println(lastName)
}