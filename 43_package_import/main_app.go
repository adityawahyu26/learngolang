// package dan import package dalam Go
// package menyimpan berbagai source code Go dan tools
package main

// import untuk memanggil package dan toolsnya
import (
	// import package 
	"learngolang/43_package_import/Helper"
	"fmt"
)

func main() {
	// memakai func sayHello() dari package Helper
	hallo := Helper.SayHello("Aditya")
	fmt.Println(hallo)

	fmt.Println(Helper.SayHello("Wahyu"))
}