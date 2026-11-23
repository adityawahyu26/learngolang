// container/ring dalam bahasa pemrograman Go
// digunakan untuk menangani struktur data circular ring
package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// membuat data ring dan menentukan sizenya
	data := ring.New(5)
	// membuat loop untuk mengisi value dalam setiap data ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.FormatInt(int64(i+1), 10)
		data = data.Next()
	}

	// menampilkan semua value yang ada dalam setiap data ring
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
