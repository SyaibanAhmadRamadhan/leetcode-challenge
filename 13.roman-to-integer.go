package leetcode

// subtractive
// •	IV = 4
// 	•	IX = 9
// 	•	XL = 40
// 	•	XC = 90
// 	•	CD = 400
// 	•	CM = 900

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start

// cara 1
// func romanToInt(s string) int {
// 	runes := []rune(s)
// 	mapping := map[rune]int{
// 		'M': 1000,
// 		'D': 500,
// 		'C': 100,
// 		'L': 50,
// 		'X': 10,
// 		'V': 5,
// 		'I': 1,
// 	}

// 	count := 0
// 	for i := 0; i < len(runes); {
// 		curr := mapping[runes[i]]
// 		if len(s) > i+1 && (curr < mapping[runes[i+1]]) {
// 			count += ((mapping[runes[i+1]]) - curr)
// 			i += 2
// 		} else {
// 			fmt.Println(curr)
// 			count += curr
// 			i += 1
// 		}
// 	}

// 	return count

// }

// func romanToInt(s string) int {
// 	var romanVal = [256]int{
// 		'I': 1, 'V': 5, 'X': 10, 'L': 50,
// 		'C': 100, 'D': 500, 'M': 1000,
// 	}

// 	total, prev := 0, 0

// 	for i := len(s) - 1; i >= 0; i-- {
// 		v := romanVal[s[i]]
// 		if v < prev {
// 			total -= v
// 		} else {
// 			total += v
// 			prev = v
// 		}
// 	}

// 	return total

// }

// func romanToInt(s string) int {
// 	var romanVal = [256]int{
// 		'I': 1, 'V': 5, 'X': 10, 'L': 50,
// 		'C': 100, 'D': 500, 'M': 1000,
// 	}

// 	total := 0

// 	for i := 0; i < len(s); i++ {
// 		v := romanVal[s[i]]
// 		next := v
// 		if i+1 < len(s) {
// 			next = romanVal[s[i+1]]
// 		}

// 		if v < next {
// 			total += (next - v)
// 			i++
// 		} else {
// 			total += v
// 		}
// 	}

// 	return total

// }

func romanToInt(s string) int {
	var romanVal = [256]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	total := 0
	n := len(s)
	i := 0

	for i < n {
		if i+1 < n {
			pair := uint16(s[i])<<8 | uint16(s[i+1])
			switch pair {
			case 'I'<<8 | 'V':
				total += 4
				i += 2
				continue
			case 'I'<<8 | 'X':
				total += 9
				i += 2
				continue
			case 'X'<<8 | 'L':
				total += 40
				i += 2
				continue
			case 'X'<<8 | 'C':
				total += 90
				i += 2
				continue
			case 'C'<<8 | 'D':
				total += 400
				i += 2
				continue
			case 'C'<<8 | 'M':
				total += 900
				i += 2
				continue
			}
		}

		total += romanVal[s[i]]
		i++
	}

	return total

}

// @lc code=end
