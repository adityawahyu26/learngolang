// function sebagai parameter bagi function lain
// membantu tugas function lainnya
package main

import "fmt"

// type declaration untuk params berbentuk func
type Filter func(string) string

// func dengan params string dan func lainnya
func wordChecker(name string, filter Filter) string {
	// menyimpan value sekaligus run func lainnya
	wordFiltered := filter(name)
	// return yang juga menerima return func lainnya
	return "hello " + wordFiltered
}

// func yang dijadikan parameter
// menerima value dari func lainnya dalam bentuk params
func filterChecker(name string) string {
	// cek conn params
	if name == "badwords" {
		// jika conn ideal false, return value dengan str
		return "***"
	} else {
		// conn ideal dimana mengembalikan value asli
		return name
	}
}

func main()  {
	// case 1, func("str",funcparams)
	word := wordChecker("aditya", filterChecker)
	fmt.Println(word)

	// case 2, menyimpan dalam variable
	person1 := "wahyu"
	filter1 := filterChecker
	func1 := wordChecker(person1, filter1)
	fmt.Println(func1)

	// case 3, langsung print func("str",funcparams)
	fmt.Println(wordChecker("badwords",filterChecker))
}