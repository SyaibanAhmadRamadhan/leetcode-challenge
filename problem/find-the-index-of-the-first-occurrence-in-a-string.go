package problem

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if haystack == needle {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if len(haystack[i:]) >= len(needle) &&
			haystack[i:len(needle)+i] == needle {
			return i
		}
	}

	return -1
}
