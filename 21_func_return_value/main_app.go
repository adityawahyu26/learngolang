// return value di function dalam bahasa pemrograman Go
// mengembalikan nilai dari func, harus di define type datanya
package main

import "fmt"

// func namafunc(params tipedata) return tipedata {
// menerima parameter name yang bertipe data string
// dan akan menghasilkan return dengan tipe data string
func sayHello(name string) string {
	// menggunakan if else untuk conn berdasarkan params
	if name == "" {
		return "hello master"
	} else {
		return "hello " + name
	}
}

func main() {
	// case dengan string kosong
	user := sayHello("")
	fmt.Println(user)

	// case dengan string 
	person := sayHello("aditya")
	fmt.Println(person)
}