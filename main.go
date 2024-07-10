package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Isvalid("1234567890123"))
}

func Isvalid(id string) bool {
	sum := 0
	multi := []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(id[i]))
		sum += multi[i] * digit
	}

	remainder := sum % 11
	remainder = 11 - remainder
	digit := remainder

	if remainder > 10 {
		for remainder > 10 {
			remainder = remainder % 10
			digit = remainder
		}
	}

	lastDigit, _ := strconv.Atoi(string(id[12]))
	return digit == lastDigit
}
