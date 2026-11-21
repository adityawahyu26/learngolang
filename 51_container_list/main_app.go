// container/list dalam bahasa pemrograman Go
// digunakan untuk strukur data double linked list
// next() untuk akses data selanjutnya
// prev() untuk akses data sebelumnya
// nil - data - data2 - data3 - nil
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// deklarasi variable
	data := list.New()

	// input datanya, pushback untuk selalu input diakhir data
	data.PushBack("aditya")
	data.PushBack("wahyu")
	data.PushBack("x")
	data.PushBack("y")

	// loop untuk menampilkan semua value datanya
	// data.Front() menampilkan data yang ada di depan
	// i != nil, selama loop i tidak boleh sampai nil
	// i next , akan maju satu data setelahnya
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
