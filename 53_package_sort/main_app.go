// package sort dalam bahasa pemrograman Go
// digunakan untuk mengurutkan data
package main

import (
	"fmt"
	"sort"
)

// deklarasi struct
type User struct {
	name string
	age  int
}

// type slice of user untuk menyimpan banyak struct
type UserSlice []User

// func untuk mengecek panjang []struct
func (value UserSlice) Len() int {
	return len(value)
}

// func untuk memastikan bahwa parameter i harus lebih kecil dari j
func (value UserSlice) Less(i, j int) bool {
	return value[i].age < value[j].age
}

// func untuk menukar posisi parameter i dan j
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	// membuat kumpulan data struct
	person := []User{
		{"aditya", 45},
		{"wahyu", 12},
		{"x", 86},
		{"y", 23},
	}

	// sebelum di sort
	fmt.Println(person)
	// untuk menggunakan func Sort di package sort harus bisa pakai Interface di sort
	// untuk bisa pakai Interface, harus punya func Len(), Less(), Swap()
	// jadi UserSlice bisa pakai func Sort karena lolos Interface
	sort.Sort(UserSlice(person))
	// data setelah di sort
	fmt.Println(person)
}
