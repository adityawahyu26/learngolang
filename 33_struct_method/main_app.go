// struct method dalam bahasa pemrograman Go
// method yang bisa menerima dan dijalankan struct
package main

import "fmt"

// deklarasi struct
type User struct {
	nickname string
	region string
}

// method stuct, menerima struct dan bisa run fieldnya
func (user User) sayHello() {
	fmt.Println("hello, my name is", user.nickname)
	fmt.Println("and i'm from", user.region)
}

func main() {
	// deklarasi variable struct
	adityawp := User{
		nickname: "aditya wahyu",
		region: "indonesia",
	}
	// memanggil struct method
	adityawp.sayHello()
}