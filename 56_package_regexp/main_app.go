// package regexp dalam bahasa pemrograman Go
// digunakan untuk menangani reguler expression dalam program
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// membuat variable regexp
	regex := regexp.MustCompile(`a([a-z])i`)

	// mencoba string match dengan regexp
	fmt.Println(regex.MatchString("abi"))
	fmt.Println(regex.MatchString("aci"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("afa"))

	// menemukan semua string yang match dengan regexp dan jumlah batas pencariannya
	// -1 artinya jumlahnya semua string yang match dengan regexp
	fmt.Println(regex.FindAllString("abi ace axu aci ake adi", -1))
}
