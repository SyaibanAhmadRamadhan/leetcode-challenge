package problem

func ReverseWords(s string) string {
	j := 0
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		if b[i] == ' ' {
			continue
		}

		if j > 0 {
			b[j] = ' '
			j++
		}

		for i < len(b) && b[i] != ' ' {
			b[j] = b[i]
			i++
			j++
		}
	}

	for i, j := 0, len(b[:j])-1; i <= j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	start := 0
	for z := 0; z <= len(b[:j])-1; z++ {
		if b[z] == ' ' {
			for i, j := start, z-1; i <= j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}
			start = z + 1
		}
		if z == len(b[:j])-1 {
			for i, j := start, z; i <= j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}
		}
	}

	return string(b[:j])
}
