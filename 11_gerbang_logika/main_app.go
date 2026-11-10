// gerbang logika dalam bahasa pemrograman Go
package main

import "fmt"

func main() {
	// menggunakan AND (&&)
	var nilai = 90
	var absensi = 80

	var lulus_nilai = nilai >= 75
	var lulus_absensi = absensi >= 75

	var lulus = lulus_nilai && lulus_absensi
	fmt.Println("Lulus Nilai:", lulus_nilai)
	fmt.Println("Lulus Absensi:", lulus_absensi)
	fmt.Println("Lulus:", lulus)

	// menggunakan OR (||)
	var tiket_cash = true
	var tiket_online = false
	var bisa_masuk = tiket_cash || tiket_online
	fmt.Println("Bisa Masuk:", bisa_masuk)

	// menggunakan NOT (!)
	var sakit = false
	fmt.Println("apa dia sakit:", sakit)
	fmt.Println("apa dia bisa sekolah:", !sakit)
}