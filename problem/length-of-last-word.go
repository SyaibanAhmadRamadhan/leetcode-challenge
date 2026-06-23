package problem

import "strings"

func LengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)

	splitString := strings.Split(s, " ")
	return len(splitString[len(splitString)-1])
}
