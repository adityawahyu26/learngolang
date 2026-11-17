// assertions dalam bahasa pemrograman Go
// mengubah satu tipe data ke tipe data lainnya
// biasanya dari tipe any / interface{} ke tipe data lainnya
package main

import "fmt"

// func yang returnya bertipe data any tapi valuenya int
func random() any {
	return 12
}

// func yang returnya bertipe data any tapi valuenya string
func revealType() any {
	return "aditya wahyu"
}

func main() {
	// case 1, mengubah tipe data hasil return func
	// menyimpan func dalam var
	randomAny := random()
	// sebelumnya return func bertipe data any
	fmt.Println("ini bertipe data any:", randomAny)
	// lalu diubah menjadi int, sesuai tipe data return valuenya
	randomInt := randomAny.(int)
	// akan berubah dari tipe data any ke tipe data int
	fmt.Println("ini bertipe data int:", randomInt)

	// case 2, cek apa tipe data sebenarnya dari return value
	reveal := revealType()
	// mengambil tipe data dari hasil return value func
	switch v := reveal.(type) {
	// hasilnya akan sesuai dengan tipe data sebenarnya value
	case string: 
		fmt.Println("nilai", v, "bertipe data string")
	case int: 
		fmt.Println("nilai", v, "bertipe data int")
	case bool: 
		fmt.Println("nilai", v, "bertipe data bool")
	default:
		fmt.Println("Unknown type")
	}	
}