// package time dalam bahasa pemrograman Go
// mengakses dan memanipulasi waktu di dalam program
package main

import (
	"fmt"
	"time"
)

func main() {
	// menampilkan waktu saat ini
	// dimulai dari tahun, bulan, tanggal, jam, menit, detik, nanodetik
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Println("\n")

	// membuat objek time
	date := time.Date(2005, 02, 26, 12, 18, 43, 21, time.Local)
	fmt.Println(date)

	fmt.Println("\n")

	// parse dari data string ke objek time
	parse, _ := time.Parse("2006-01-02", "2005-02-26")
	fmt.Println(parse)
}
