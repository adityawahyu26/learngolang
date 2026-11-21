// package initialization dalam Go
// akan selalu dijalankan saat package diakses
package main

import (
	// import package
	"learngolang/45_package_initialization/Database"
	"fmt"
	// import package dengan blank import
	// untuk hanya menggunakan initnya tanpa func
	_"learngolang/45_package_initialization/BackupDB"
)

func main() {
	// memanggil value init dengan func dari package
	result := Database.GetDatabase()
	fmt.Println(result)
}