// variadic function dalam bahasa pemrograman Go
// menyimpan parameter/argument yang punya value lebih banyak

package main

import "fmt"
// menerima parameter number yang bisa berisi banyak value
func sum(number ...int) int{
	// membuat func operasi tambah sederhana
	total := 0
	// mengambil setiap _(index) dan value dari number
	for _,value := range number {
		total += value
	}
	return total
}

func main() {
	// memanggil func dan input parameter yang berisi beberapa value
	fmt.Println(sum(12,13,14,15))

	numbers := []int{1,2,3,4,5}
	// menggunakan slice untuk input parameternya
	fmt.Println(sum(numbers...))
}