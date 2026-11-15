// function value dalam Go
package main

import "fmt"

// func hello 
func sayHello(name string) {
	if name == "" {
		fmt.Println("hello master")
		fmt.Println("have a nice day!")
	} else {
		fmt.Println("hello", name)
		fmt.Println("have a nice day!")
	}
}

func main() {
	// memanggil func secara biasa
	sayHello("")

	// menyimpan func dalam sebuah variable
	hello := sayHello
	// menggunakan variable untuk menjalankan func
	hello("aditya wahyu")
}