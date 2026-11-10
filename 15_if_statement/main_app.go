// if statement dalam bahasa pemrograman Go
// if statement akan dieksekusi jika conn tertentu true / terpenuhi
package main

import "fmt"

func main() {
	name := "aditya"

	// if statement, akan dieksekusi jika conn tertentu true
	if name == "aditya" {
		fmt.Println("hallo, ", name)
	}

	// akan dieksekusi selama conn tertentu terpenuhi
	// else adalah alternatif jika conn bernilai false
	rice_cooker := "warm"
	if rice_cooker == "warm" {
		fmt.Println("nasi sudah matang dan siap disantap")
	} else {
		fmt.Println("nasi belum matang, tunggu beberapa saat!")
	}

	// menggunakan if dan else if statement
	// else if adalah case cabang untuk conn true lainnya
	number := 23 + 32 - 12 * 123 / 293
	fmt.Println("[ 23 + 32 - 12 * 123 / 293 ]")
	if number > 100 {
		fmt.Println("hasil perhitungan > 100:", number)
	} else if number < 100 {
		fmt.Println("hasil perhitungan < 100:", number)
	} else {
		fmt.Println("hasil perhitungan tidak valid")
	}

	// membuat var dalam if statement
	num := [5]int{21,22,23,24,25}
	fmt.Println(num)
	if length := len(num); length > 5 {
		fmt.Println("terdapat lebih dari 5 value didalamnya")
	} else {
		fmt.Println("terdapat 5 value atau dibawahnya")
	}
}