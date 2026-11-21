// package math dalam bahasa pemrograman Go
// digunakan untuk operasi matematika dalam program
package main

import (
	"fmt"
	"math"
)

func main() {
	// membulatkan ke yang terdekat, bisa bawah atau atas
	fmt.Println("round:", (math.Round(12.7)))
	fmt.Println("round:", (math.Round(12.3)))

	// selalu membulatkan ke bawah
	fmt.Println("floor:", (math.Floor(12.7)))
	fmt.Println("floor:", (math.Floor(12.3)))

	// selalu membulatkan ke atas
	fmt.Println("ceil:", (math.Ceil(12.7)))
	fmt.Println("ceil:", (math.Ceil(12.3)))

	// nilai paling banyak dari dua angka yang disandingkan
	fmt.Println("max (12/26):", (math.Max(12, 26)))

	// nilai paling sedikit dari dua angka yang disandingkan
	fmt.Println("min (12/26):", (math.Min(12, 26)))
}
