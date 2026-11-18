// pointer dalam function dalam Go
// mengakses dan mengubah value dari pointer value
package main

import "fmt"

// func yang akan menerima pointer bertipe string
func status(v *string) {
	// melakukan if stmt berdasarkan params pointer
	if *v == "cook" {
		// lalu mengubah nilainya
		*v = "warm"
	} else if *v == "warm" {
		*v = "cook"
	} else {
		*v = "unknown type"
	}
}

// func yang menerima pointer bertipe boolean
func saklar(value *bool) {
	// mengembalikan dalam bentuk kebalikan dari params
	*value = !(*value)
}

func main() {
	// case 1, mengubah status rice cooker
	// setiap kali func dipanggil
	rice_cooker := "cook"
	kitchen_tools := &rice_cooker
	// memasukkan pointer kedalam func sebagai params
	status(kitchen_tools)
	// lalu cek, apa variable awal dan pointer berubah
	fmt.Println("kondisi rice cooker:", rice_cooker)
	// mencoba melakukan operasi sebaliknya
	status(kitchen_tools)
	// hasilnya akan sebaliknya juga 
	fmt.Println("kondisi rice cooker", *kitchen_tools)

	// membuat struct 
	type Lampu struct {
		status bool
	}

	// membuat variable struct 
	lampu_kamar := Lampu{false}
	// mengirim salah satu field yang bertipe boolean
	saklarLampu := &lampu_kamar.status
	// kirim ke func
	saklar(saklarLampu)
	// maka hasil nya adalah boolean yang sudah reverse
	fmt.Println("status lampu kamar:", *saklarLampu)

	// melakukan operasi sebaliknya 
	// hasilnya akan selalu berlawanan
	lampu_ruang_tamu := Lampu{true}
	saklarLampu = &lampu_ruang_tamu.status
	saklar(saklarLampu)
	fmt.Println("status lampu ruang tamu:", *saklarLampu)
}