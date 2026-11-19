package BackupDB

import "fmt"

var database string

// innit akan selalu di run begitu package ini diakses
func init() {
	fmt.Println("innit dipanggil, (ini dari BackupDB/backup.go)")
	database = "MySQL"
}

// karena dipanggil dengan blank import di package lain
// maka func ini tidak akan dijalankan
func GetDatabase() string {
	return database
}