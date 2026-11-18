// comment dalam bahasa pemrograman Go
// comment yang baik adalah dalam kode itu yang sendiri 
package main

import "fmt"

/** ini adalah
	comment multi line */
func hello(name string) {
	fmt.Println("hello", name)
} 

// ini adalah comment single line
func main() {
	user := "aditya wahyu"
	hello(user)
}