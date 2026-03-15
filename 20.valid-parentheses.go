package leetcode

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) <= 32 {
		stack := uint64(0)
		trash := 0

		first := func(s rune) uint64 {
			switch s {
			case '(':
				return 1
			case '{':
				return 2
			case '[':
				return 3
			default:
				return 0
			}
		}

		matchLatest := func(s rune) uint64 {
			switch s {
			case ')':
				return 1
			case '}':
				return 2
			case ']':
				return 3
			default:
				return 0
			}
		}

		for _, v := range s {
			if code := first(v); code != 0 {
				stack = (stack << 2) | code
				trash++
				continue
			}

			if code := matchLatest(v); code != 0 {
				if trash == 0 {
					return false
				}

				latest := stack & 0b11
				if latest != code {
					return false
				}

				stack >>= 2
				trash--
				continue
			}

			return false
		}

		return trash == 0
	} else {
		stack := make([]rune, 0, len(s))
		for _, v := range s {
			switch v {
			case '(', '{', '[':
				stack = append(stack, v)
			case ')', '}', ']':
				if len(stack) == 0 {
					return false
				}

				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if (v == ')' && top != '(') ||
					(v == '}' && top != '{') ||
					(v == ']' && top != '[') {
					return false
				}
			default:
				return false
			}
		}

		return len(stack) == 0
	}
}

// @lc code=end
