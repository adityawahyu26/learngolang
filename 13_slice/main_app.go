// slice dalam bahasa pemrograman Go
// digunakan untuk mengambil bagian tertentu dari sebuah array
package main

import "fmt"

func main() {
	// deklarasi variable
	var month = [...]string{
		"January", 
		"February", 
		"March", 
		"April", 
		"May", 
		"June", 
		"July", 
		"August", 
		"September", 
		"October", 
		"November", 
		"December"}

	fmt.Println(month)

	// membuat slice dari array
	// [2:5] artinya dimulai dari index 2 (march) - index 4 (may) 
	slice1 := month[2:5]
	fmt.Println(slice1)
	// panjang dari value yang ditampung slice, yaitu 3 value
	fmt.Println(len(slice1))
	// kapasitas slice, dimulai dari index slice awal - index terakhir
	fmt.Println(cap(slice1))
	
	// membuat slice dengan index awal 2 - index terakhir array
	slice2 := month[2:]
	fmt.Println(slice2)
	// len datanya dihitung dari index awal slice - index terakhir
	fmt.Println(len(slice2))
	// pada metode ini sama dengan len, index awal - index terakhir
	fmt.Println(cap(slice2))

	// slice yang dimulai dari index 0 - index 5 array
	slice3 := month[:5]
	fmt.Println(slice3)
	// len slice di hitung dari index 0 sampai index 5
	fmt.Println(len(slice3))
	// untuk kapasitas dihitung dari index awal (0) - index akhir
	fmt.Println(cap(slice3))

	// membuat slice yang mengambil dari index awal - index akhir
	slice4 := month[:]
	fmt.Println(slice4)
	// len slice ini dari index awal hingga index terakhir array
	fmt.Println(len(slice4))
	// kapasitas slice ini dimulai dari index awal - index terakhir
	fmt.Println(cap(slice4))

	// mengubah slice akan mempengaruhi array referencenya
	// dimana array dan slice lain datanya juga bisa berubah
	slice4[0] = "awal bulan"
	// fmt.Println(month)
	slice4[0] = "january"

	// membuat slice baru dengan func make() tanpa array reference
	// slice = make([]type data, len slice, cap slice)
	var number = make([]int, 3, 5)
	// insert value dalam slice
	number[0] = 12
	number[1] = 26
	number[2] = 44
	fmt.Println(number)

	// membuat array dengan 5 buah value
	no_antri := [5]int{1,2,3,4,5}
	// buat slice untuk arraynya
	// dimulai dari index ke-3 sampai index akhir
	slice5 := no_antri[3:]
	fmt.Println(slice5)
	// append data baru pada slicenya
	// jika kapasitas array penuh, akan bikin array baru
	new_no_antri := append(slice5, 6)
	// array baru berisi value dari slicenya dan data yang di add
	// bisa diubah datanya untuk array baru
	fmt.Println(new_no_antri)
	fmt.Println(no_antri)

	// membuat array
	nomor := [5]int{11,12,13,14,15}
	// memasukkan semua value array ke dalam slice
	slice6 := nomor[:]
	// buat struktur slice baru yang len dan cap reference slice lama
	copy_nomer := make([]int, len(slice6), cap(slice6))
	// menggunakan func copy()
	copy(copy_nomer, slice6)
	// copy slice tidak reference dengan slice lama
	// walau value dan struktur sama persis, tapi keduanya berbeda
	// jadi bisa diubah tanpa ganggu yang lainnya
	fmt.Println(nomor)
	fmt.Println(copy_nomer)
}