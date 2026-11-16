// interface dalam bahasa pemrograman Go
// menyimpan beberapa method struct dalam satu tempat
package main

import "fmt"

// deklarasi untuk type struct southeast asia
type SoutheastAsia struct {
	name string
}

// membuat method yang menerima struct southeast asia
func (country SoutheastAsia) asiaUnity() {
	// membuat infografis sederhana di method
	fmt.Println(country.name, "is country in the southeast asia")
}

// deklarasi untuk type struct east asia
type EastAsia struct {
	name string
}

// membuat method yang menerima struct east asia
func (country EastAsia) asiaUnity() {
	// membuat infografis sederhana di method
	fmt.Println(country.name, "is country in the east asia")
}

// membuat interface yang berisi persyaratan jika ingin dipakai
type AsiaRegion interface {
	// setiap structnya harus punya method asiaUnity()
	asiaUnity()
}

// method menerima struct yang memenuhi syarat interface asia
func Asia(country AsiaRegion) {
	// memanggil method asiaUnity()
	// akan ditampilkan sesuai jenis type structnya
	country.asiaUnity()
}

func main() {
	// deklarasi variable bertype struct
	Indonesia := SoutheastAsia{
		name: "Indonesia",
	}

	Japan := EastAsia{
		name: "Japan",
	}

	// run method asia, dan akan dicek apa lulus interface asia
	// jika lulus, akan ditampilkan sesuai jenis type structnya
	Asia(Indonesia)
	Asia(Japan)
}