package main

import (
	"fmt"
)

func sumDigits(number int) int {
	remainder := 0
	sumResult := 0
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	if sumResult > 9 {
		return sumDigits(sumResult)
	}
	return sumResult
}

func main() {
	var number int
	fmt.Print("Enter Number:")
	fmt.Scanln(&number)
	fmt.Printf("Addition digits of %d = %d\n", number, sumDigits(number))
}
