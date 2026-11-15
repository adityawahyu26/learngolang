// defer, panic dan recover dalam Go
package main

import "fmt"

func root() {
	fmt.Println("hello master")
	fmt.Println("have a nice day")
	// mengecek apa ada runtime error dan run recover
	if r:= recover(); r != nil{
		fmt.Println("program berhasil di recover", r)
	}
}

func runApp(value int) {
	// defer akan selalu run walau func runApp runtime error 
	defer root()
	i := value
	fmt.Println("\nrunning...")
	// panic di eksekusi untuk mencegah runtime error 
	if i == 0 {
		panic("\nangka 10 tidak bisa dibagi 0")
	}
	// jika runtime error, efek panic() akan cegah ini muncul
	fmt.Println(10/i)
}

func main() {
	// run case jika ada runtime error
	runApp(0)
	// run case jika tidak ada error apapun
	runApp(2)
}