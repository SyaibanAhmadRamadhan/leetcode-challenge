package leetcode

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	s := "()[]{}"

	fn := func() bool {
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
				fmt.Println(string(top))
				fmt.Println(string(v))
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

	fmt.Println(fn())
}
