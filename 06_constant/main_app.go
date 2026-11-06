// constant digunakan untuk menyimpan value yang tetap
// const harus diinisiasi sekali dan tidak bisa diubah valuenya

package main

import "fmt"

func main() {
	const (
		firstName = "Aditya"
		lastName = "Wahyu"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	const pi float64 = 3.14
	fmt.Println(pi)
}