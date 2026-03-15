package leetcode

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	enableLogic := 4

	if enableLogic == 1 {
		runes := []rune(s)
		mapping := map[rune]int{
			'M': 1000,
			'D': 500,
			'C': 100,
			'L': 50,
			'X': 10,
			'V': 5,
			'I': 1,
		}

		count := 0
		for i := 0; i < len(runes); {
			curr := mapping[runes[i]]
			if len(runes) > i+1 && curr < mapping[runes[i+1]] {
				curr = mapping[runes[i+1]] - curr
				count += curr
				i += 2
			} else {
				count += curr
				i++
			}
		}

		return count
	} else if enableLogic == 2 {
		romanVal := [128]int16{
			'M': 1000,
			'D': 500,
			'C': 100,
			'L': 50,
			'X': 10,
			'V': 5,
			'I': 1,
		}

		total, prevValue := int64(0), int16(0)
		for i := len(s) - 1; i >= 0; i-- {
			value := romanVal[s[i]]
			if value < prevValue {
				total -= int64(value)
			} else {
				total += int64(value)
				prevValue = value
			}
		}

		return int(total)
	} else if enableLogic == 3 {
		romanVal := [128]int16{
			'M': 1000,
			'D': 500,
			'C': 100,
			'L': 50,
			'X': 10,
			'V': 5,
			'I': 1,
		}

		total := int64(0)
		for i := 0; i < len(s); i++ {
			value := romanVal[s[i]]
			next := value
			if len(s) > i+1 {
				next = romanVal[s[i+1]]
			}

			if value < next {
				total += (int64(next) - int64(value))
				i++
			} else {
				total += int64(value)
			}
		}

		return int(total)
	} else if enableLogic == 4 {
		// bit shifting used
		romanVal := [256]int{
			'I': 1, 'V': 5, 'X': 10, 'L': 50,
			'C': 100, 'D': 500, 'M': 1000,
		}

		count := 0
		n := len(s)
		i := 0

		for i < n {
			if i+1 < n {
				pair := uint16(s[i])<<8 | uint16(s[i+1])
				switch pair {
				case 'I'<<8 | 'V':
					count += 4
					i += 2
					continue
				case 'I'<<8 | 'X':
					count += 9
					i += 2
					continue
				case 'X'<<8 | 'L':
					count += 40
					i += 2
					continue
				case 'X'<<8 | 'C':
					count += 90
					i += 2
					continue
				case 'C'<<8 | 'D':
					count += 400
					i += 2
					continue
				case 'C'<<8 | 'M':
					count += 900
					i += 2
					continue
				}
			}

			count += romanVal[s[i]]
			i++
		}
		return count
	} else {
		return 1
	}
}

// @lc code=end
