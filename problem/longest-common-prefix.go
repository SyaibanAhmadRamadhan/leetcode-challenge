package problem

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	prefix := strs[0]
	for _, s := range strs[1:] {
		for !strings.HasPrefix(s, prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	// for i := 1; i < len(strs); i++ {
	// 	for len(prefix) > 0 {
	// 		if len(strs[i]) >= len(prefix) &&
	// 			strs[i][:len(prefix)] == prefix {
	// 			break
	// 		}

	// 		prefix = prefix[:len(prefix)-1]
	// 	}
	// 	if prefix == "" {
	// 		return ""
	// 	}
	// }

	return prefix
}
