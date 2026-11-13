// multiple return values di function dalam bahasa pemrograman Go
// mengembalikan beberapa value sekaligus dari function
package main

import "fmt"

// case dua value
// func namafunc() (tipedata return,..) {
func footballClub() (string, string) {
	// deklarasi dua var
	club := "liverpool"
	league := "premier league"

	// return keduanya
	return club, league
}

func main() {
	// var pertama menyimpan return value pertama
	// var kedua menyimpan return value kedua
	club, league := footballClub()
	// tampilkan outputnya
	fmt.Println(club)
	fmt.Println(league)
}