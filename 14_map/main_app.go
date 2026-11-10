// map dalam bahasa pemrograman Go
// mirip array assosiatif, menyimpan array dengan index custom
// untuk index/key nya bisa diubah sendiri
package main

import "fmt"

func main() {
	// deklarasi map dengan key dan value bertipe data string
	person := map[string]string {
		"name":"aditya wahyu",
		"country":"indonesia",
	}
	// menampilkan map, bisa dipanggil berdasarkan key nya
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])
	// len(map), digunakan untuk mengecek len dari sebuah map
	fmt.Println("panjang sebelum ditambahkan data:", len(person))
	
	// menambahkan value baru pada map
	person["food"] = "mie ayam"
	fmt.Println(person)
	// mengecek len baru setelah value yang baru di add ke map
	fmt.Println("panjang setelah ditambahkan data:", len(person))

	// membuat map dengan func make()
	indonesia := make(map[string]string) 
	// mengisi value kedalam map
	indonesia["capital"] = "jakarta"
	indonesia["currency"] = "rupiah"
	indonesia["animal"] = "komodo"
	indonesia["error"] = "..."
	
	fmt.Println(indonesia)
	
	// menghapus value tertentu berdasarkan key nya di map
	delete(indonesia, "error")
	// mengecek apa perubahan di eksekusi program
	fmt.Println(indonesia)
}