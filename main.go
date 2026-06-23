package main

import (
	"fmt"
	"leetcode/problem"
)

func main() {
	output := problem.FullJustify([]string{
		"This", "is", "an", "example", "of", "text", "justification.",
	}, 16)
	for _, v := range output {
		fmt.Println(v)
	}
}
