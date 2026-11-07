// konvers type data digunakan untuk mengubah tipe data
// bisa dari yang berbeda jenis seperti string, int, float, boolean
// atau yang sama jenis seperti int8, int16, int32, int64
package main

import "fmt"

func main() {
	var int32 int32 = 128
	var int64 int64 = int64(int32)
	var int8 int8 = int8(int32)
	
	fmt.Println("int32:", int32)
	fmt.Println("int64:", int64)
	fmt.Println("int8:", int8)
}