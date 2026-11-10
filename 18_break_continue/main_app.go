// break dan continue dalam bahasa pemrograman Go
// digunakan untuk menghentikan / melanjutkan loop
package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {
		// break akan menghentikan loop pada perulangan saat i == 12
		// perulangan selanjutnya tidak akan dilanjutkan 
		// dan program berhenti pada conn itu
		if i == 12 {
			fmt.Println("We'll be back soon")
			break
		}
		fmt.Println("Loop =>", i)
	}

	for i := 1; i <= 20; i++ {
		// continue akan melanjutkan loop berdasarkan conn tertentu
		// case bil genap, akan mengabaikan saat loop i == bil genap
		// dan lanjut ke loop berikutnya
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Ganjil >>", i)
	}
	
	for i := 1; i <= 20; i++ {
		// case untuk bilangan ganjil / i == bil ganjil
		if i % 2 == 1 {
			continue
		}
		fmt.Println("Genap >>", i)
	}
}