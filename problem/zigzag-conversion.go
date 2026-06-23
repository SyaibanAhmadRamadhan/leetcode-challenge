package problem

import "strings"

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := make([][]byte, numRows)
	rowCurrent := 0
	dir := 1
	for i := 0; i < len(s); i++ {
		if rowCurrent == numRows-1 {
			dir = -1
		}
		arr[rowCurrent] = append(arr[rowCurrent], s[i])
		rowCurrent += dir
		if rowCurrent == 0 {
			dir = 1
		}
	}

	var str strings.Builder
	for _, v := range arr {
		str.Write(v)
		println(string(v))
	}
	return str.String()
}
