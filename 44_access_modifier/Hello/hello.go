package Hello

// func ini bisa dipanggil dari package lain
// karena huruf depannya adalah kapital
func GoodMorning(name string) string {
	return "hello, good morning " + name
}

// func ini akan gagal dipanggil di func lain
// dan hanya bisa dipakai di package ini
// karena huruf depan namanya bukan kapital
func goodNight(name string) string {
	return "hello, good night " + name
}