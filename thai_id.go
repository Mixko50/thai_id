package main

import (
	"strconv"
)

func thaiId(id string) bool {
	if len(id) != 13 {
		return false
	}

	sumAllDigitResult := 0
	lastDigit, _ := strconv.Atoi(string(id[12]))

	for i := 0; i < 12; i++ {
		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		sumAllDigitResult += num * (13 - i)
	}

	remainder := sumAllDigitResult % 11
	minusResult := 11 - remainder
	remainder = minusResult % 10

	if remainder != lastDigit {
		return false
	}

	return true
}
