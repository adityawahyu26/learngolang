// nill dalam bahasa pemrograman Go
// nill tidak berisi value atau nilai apapun dan hanya kosong
package main

import "fmt"

// func cek apa user sudah login dan return value map
func login(user string) map[string]string {
	// jika parameter kosong akan return nil
	if user == ""{
		return nil
	// jika parameter ada isinya akan return dan buat map
	} else {
		return map[string]string{
			"name": user,
		}
	}
}

func main() {
	// deklarasi var user
	user := ""
	// user := "aditya"
	
	// memanggil func dengan var
	userLogin := login(user)

	// jika hasil return func nil akan menampilkan ini
	if userLogin == nil {
		fmt.Println("silahkan login terlebih dahulu")
	// jika return ada isinya maka akan menampilkan mapnya
	} else {
		fmt.Println("selamat datang", userLogin["name"])
	}
}