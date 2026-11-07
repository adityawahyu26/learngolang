package main

import "fmt"

func main() {
	// mendeklarasikan type data int dengan nama noPesanan
	type noPesanan int

	// jadi tidak perlu var pesanan_andi int = 100
	// cukup var pesanan_andi noPesanan = 100
	var pesanan_andi noPesanan = 100
	fmt.Println(pesanan_andi)

	var pesanan_budi noPesanan = 200
	fmt.Println(pesanan_budi)
}