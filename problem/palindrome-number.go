package problem

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	divisor := 1
	for x/divisor >= 10 {
		divisor *= 10
	}

	for x > 0 {
		first := x / divisor
		last := x % 10

		if first != last {
			return false
		}

		x = x % divisor / 10
		divisor /= 100
	}

	return true
}
