func isPowerOfThree(n int) bool {
	if n < 3 {
		if n == 1 {
			return true
		} else {
			return false
		}
	}
	if n%3 != 0 {
		return false
	}
	if n == 1 {
		return true
	}
	return isPowerOfThree(n / 3)
}