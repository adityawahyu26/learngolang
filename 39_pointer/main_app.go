// pointer dalam bahasa pemrograman Go
// digunakan untuk mengubah atau mengakses value
// dari tempat lain yang merupakan pointernya
package main

import "fmt"

func main() {
	// deklarasi struct
	type structBiasa struct {
		nama, kota string
	}

	type Address struct {
		city, province, country string
	}

	// metode copy biasa
	varBiasa1 := structBiasa {
		nama: "aditya",
		kota: "surakarta",
	}
	varBiasa2 := varBiasa1
	// mengubah value salah satu copy an
	varBiasa2.nama = "wahyu"
	// variable awalnya tidak ikut berubah
	fmt.Println(varBiasa1)
	// hanya variable yang diubah yang akan berubah
	fmt.Println(varBiasa2)

	// menggunakan metode pointer
	red := Address{
		city: "surabaya",
		province: "jawa timur",
		country: "indonesia",
	}
	blue := &red
	// mengakses variable awal dari pointernya
	fmt.Println(blue) 
	fmt.Println(*blue)
	// mengubah value dari dalam pointernya
	blue.city = "bandung"
	(*blue).province = "jawa barat"
	// hasilnya baik yang variable awal dan pointernya
	// akan ikut berubah 
	fmt.Println(red)
	fmt.Println(blue)

	// membuat pointer dari pointer
	yellow := &blue
	// mengakses variable sebelumnya
	fmt.Println(yellow)
	fmt.Println(*yellow)

	// mengubah value dari pointer turunan
	(*yellow).city = "semarang"
	(**yellow).province = "jawa tengah"
	// ini semua akan menghasilkan output yang sama
	// variable awal
	fmt.Println(red)
	// *blue mengakses langsung isi dari red
	fmt.Println(*blue)
	// **yellow mengakses langsung isi dari red
	fmt.Println(**yellow)
}