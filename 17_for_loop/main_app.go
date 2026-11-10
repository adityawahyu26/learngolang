// for loop dalam bahasa pemrograman Go
// untuk loop yang sudah diketahui jumlah perulangannya
package main

import "fmt"

func main() {
	// deklarasi variable
	counter := 1

	// for loop untuk 10 data
	for counter <= 10 {
		fmt.Println("ini adalah perulangan ke-", counter)
		counter++
	}

	// for loop dengan short statement, mengulang 5 data
	for i := 11; i <= 15; i++ {
		fmt.Println("loop =>", i)
	} 
	
	// membuat array untuk loop
	number := [5]int{12,34,53,21,34}

	// loop untuk setiap value array berdasarkan len array
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}

	// loop untuk setiap index dan value array
	for index,value := range number {
		fmt.Println("index =", index, ", value =", value)
	}
}