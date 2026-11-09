// switch case dalam bahasa pemrograman Go
// jika case tidak terlalu rumit pakai switch case daripada if
package main

import "fmt"

func main() {
	// deklarasi variable
	name := "wahyu"

	// switch untuk var
	switch name {
	// mencocokkan value berdasarkan case yang sesuai
	case "aditya":
		fmt.Println("welcome", name)
	case "wahyu":
		fmt.Println("welcome", name)
	// case default jika value tidak cocok di case manapun
	default:
		fmt.Println("welcome admin")
	}
	
	// switch case dengan short statement 
	// mengecek apa len > 5, true / false
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("value terlalu panjang")
	case false:
		fmt.Println("panjang value sempurna")
	}

	// deklarasi var int
	random_number := 26
	// tanpa menggunakan conn
	switch {
	// mengecek berdasar case yang sesuai
	case random_number > 30:
		fmt.Println("value ada diatas angka 30")
	case random_number > 20:
		fmt.Println("value ada diatas angka 20")
	case random_number > 10:
		fmt.Println("value ada diatas angka 10")
	default:
		fmt.Println("angka tidak valid")
	}
}