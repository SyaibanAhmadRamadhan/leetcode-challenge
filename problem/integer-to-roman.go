package problem

import (
	"strings"
)

func IntToRoman(num int) string {

	values := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	symbols := []string{
		"M",  // 1000
		"CM", // 900
		"D",  // 500
		"CD", // 400
		"C",  // 100
		"XC", // 90
		"L",  // 50
		"XL", // 40
		"X",  // 10
		"IX", // 9
		"V",  // 5
		"IV", // 4
		"I",  // 1
	}

	var s strings.Builder
	for i, v := range values {
		for {
			if num < v {
				break
			}
			num -= v
			s.WriteString(symbols[i])
		}
	}

	return s.String()
}
