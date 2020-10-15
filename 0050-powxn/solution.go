func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := pow(x, n/2)
	if n%2 == 0 {
		return res * res
	}
	return res * res * x
}