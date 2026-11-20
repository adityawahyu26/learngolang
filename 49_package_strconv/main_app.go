// package strconv dalam bahasa pemrograman Go
// mengubah dari satu tipe data ke tipe data lainnya
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// mengubah string ke boolean dengan strconv ParseBool
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println("boolean:", boolean)
	} else {
		fmt.Println("error:", err.Error())
	}

	// mengubah string ke integer dengan strconv ParseInt
	// ParseInt("string", base int, bitSize seperti int32 atau int64)
	// base int (2 Biner, 8 Oktal, 10 Desimal, 16 Heksadesimal)
	integer, e := strconv.ParseInt("1000000", 10, 64)
	if e == nil {
		fmt.Println("integer:", integer)
	} else {
		fmt.Println("error:", e.Error())
	}

	// mengubah string kedalam integer dengan Atoi
	// tidak perlu mendefisinikan base int dan bitSizenya
	value, er := strconv.Atoi("2000000")
	if er == nil {
		fmt.Println("value:", value)
	} else {
		fmt.Println("error:", er.Error())
	}
}
