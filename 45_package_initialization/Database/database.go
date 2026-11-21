package Database

// deklarasi variable
var database string

// akan dirun begitu package diakses
func init() {
	// menyimpan value dalam init
	database = "MySQL (ini dari Database/database.go)"
}

func GetDatabase() string {
	// mengembalikan value yang disimpan init
	return database
}