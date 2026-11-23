// package reflect dalam bahasa pemrograman Go
// digunakan mengakses dan memanipulasi nilai tertentu saat runtime
package main

import (
	"fmt"
	"reflect"
)

// deklarasi struct
type User struct {
	// menambahkan tag dengan key "validate" dan value "required" di field
	Name    string `validate:"required"`
	Age     int
	Address string `validate:"required"`
}

// func untuk validate, menerima parameter any dan return bool
func validate(data any) bool {
	// menggunakan reflect typeof untuk parameter data
	value := reflect.TypeOf(data)
	// melakukan loop sejumlah field yang ada di reflect
	for i := 0; i < value.NumField(); i++ {
		// mengakses setiap field
		field := value.Field(i)
		// mengecek apa setiap field punya tag key "validate" dan value "required"
		if field.Tag.Get("validate") == "required" {
			// jika ada field yang kosong atau tidak diisi akan return false
			if reflect.ValueOf(data).Field(i).Interface() == "" {
				return false
			}
		}
	}
	// jika field yang punya tag required terisi atau fieldnya tidak ada tag requirednya
	// maka return adalah true
	return true
}

func main() {
	// deklarasi variable struct
	aditya := User{
		Name:    "aditya",
		Age:     26,
		Address: "indonesia",
	}

	// mengakses struct variable dengan reflect
	usertype := reflect.TypeOf(aditya)
	// mengakses field pertama dari reflect yang berisi struct variable
	userField := usertype.Field(0)
	// mengecek jumlah keseluruhan field yang ada dalam reflect struct variable
	fmt.Println(usertype.NumField())
	// mengecek nama key dari salah satu field di reflect struct variable
	fmt.Println(userField.Name)

	// memanggil func validate untuk struct variable
	result := validate(aditya)
	fmt.Println("the result of validate aditya is", result)

	// jika salah satu field yang berisi tag required tidak diisi
	wahyu := User{
		Name: "wahyu",
		Age:  22,
		// Address: "indonesia",
	}
	fmt.Println("the result of validate wahyu is", validate(wahyu))
}
