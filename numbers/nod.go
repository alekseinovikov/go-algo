package numbers

func FindNod(a int, b int) int {
	if a > b {
		return findNod(a, b)
	}

	return findNod(b, a)
}

func findNod(a int, b int) int {
	for b != 0 {
		rest := a % b
		a = b
		b = rest
	}

	return a
}
