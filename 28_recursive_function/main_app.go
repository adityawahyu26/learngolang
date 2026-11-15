// recursive function dalam Go
// case factorial, kali bil dengan bil sebelumnya hingga 1
// contoh (3) == 3 * 2 * 1 = 6
package main

import "fmt"

// metode recursive dengan loop
func recursiveLoop(number int) int{
	result := 1
	for i:= number; i > 0; i-- {
		result *= i
	}
	return result
}

// metode recursive dengan function itu sendiri
func factorialRecursive(number int) int{
	// conn dimana loop berhenti
	if number == 1 {
		return 1
	// memicu loop dengan decrement hingga conn terpenuhi
	} else {
		return number * factorialRecursive(number-1)
	}
}

func main() {
	// menampilkan output dari nilai factorial 10
	fmt.Println(recursiveLoop(10))
	fmt.Println(10*9*8*7*6*5*4*3*2*1)
	fmt.Println(factorialRecursive(10))
}