// anonymous dalam function di bahasa pemrograman Go
// membuat langsung function dalam variable tanpa diberi nama
package main

import "fmt"

// type declaration untuk return bool dari anonymous func
type Blacklist func(string) bool

// func yang menerima params str dan func(str) bool
func registerUser(name string, blacklist Blacklist) {
	// jika name str dimasukkan func blacklist == true
	if blacklist(name) {
		fmt.Println("you are blocked", name)
		// jika name str dimasukkan func blacklist == false
		} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	// anonymous func yang menerima params str dan return bool
	isblock := func(name string) bool {
		// akan return true jika name == "scammer"
		return name == "scammer"
	} 

	// mengecek apa user itu masuk dalam list blocked user
	registerUser("aditya",isblock)
	registerUser("wahyu",isblock)
	registerUser("scammer",isblock)
}