package main

import (
	"fmt"
	"reflect"
	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")

}

func main() {

	var a int
	fmt.Println("enter a valid number:")
	fmt.Scanln(&a)
	sum := a
	var digits []int

	for a > 0 {
		digits = append(digits, a%10)
		a = a / 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {

		digits[i], digits[j] = digits[j], digits[i]
	}

	for sum/10 > 0 {
		sum = sum/10 + sum%10
	}

	fmt.Println("digits:", digits)
	fmt.Println("sum:", sum)

	x := arrayToString(digits, "")

	fmt.Println("converted :", arrayToString(digits, ""))

	fmt.Println("converted Type:", reflect.TypeOf(x))
}
