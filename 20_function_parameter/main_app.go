// paramater dalam function pada bahasa pemrograman Go
// digunakan untuk memasukan value/parameter tertentu ke func
package main

import "fmt"

// func(params 1, params 2)
func divineTrial(name string, title string) {
	// menggunakan params untuk melakukan if expression
	if title == "descender" {
		fmt.Println("the who can write the world")
		fmt.Println("welcome to this world, descender")
		fmt.Println(name, "the", title)
	} else {
		fmt.Println("hello", name,"the",title)
		fmt.Println("pass the trial to prove your nation")
	}
}

func main() {
	// untuk memanggil func harus juga menyertakan value/paramsnya
	// tidak boleh kurang atau lebih untuk jumlah paramsnya
	// sesuaikan dengan func asalnya
	divineTrial("venti","anemo archon")
	divineTrial("zhongli","geo anemo")
	divineTrial("raiden ei","electro archon")
	divineTrial("nahida","dendro archon")
	divineTrial("furina","hydro archon")
	divineTrial("mavuika","pyro archon")
	divineTrial("anastasya","cryo archon")
	divineTrial("aether","descender")
}