package problem

import "strings"

func FullJustify(words []string, maxWidth int) []string {
	currentLine := make([]string, 0)
	currentLength := 0
	result := make([]string, 0)
	for i := range words {
		if currentLength+len(words[i])+len(currentLine) <= maxWidth {
			currentLength += len(words[i])
			currentLine = append(currentLine, words[i])
		} else {
			var s strings.Builder
			maxspace := maxWidth - currentLength
			if len(currentLine) == 1 {
				s.WriteString(currentLine[0])
				for range maxspace {
					s.WriteString(" ")
				}
			} else {
				space := maxspace / (len(currentLine) - 1)
				extraspace := maxspace % (len(currentLine) - 1)
				for i, v := range currentLine {
					s.WriteString(v)
					if i != len(currentLine)-1 {
						for range space {
							s.WriteString(" ")
						}
						if i < extraspace {
							s.WriteString(" ")
						}
					}

				}
			}
			result = append(result, s.String())
			currentLength = len(words[i])
			currentLine = []string{words[i]}
		}
	}

	if len(currentLine) > 0 {
		var s strings.Builder
		diff := maxWidth - currentLength
		for i, v := range currentLine {
			s.WriteString(v)
			if i != len(currentLine)-1 {
				s.WriteString(" ")
				diff--
			}
		}
		for range diff {
			s.WriteString(" ")
		}
		result = append(result, s.String())
	}

	return result
}
