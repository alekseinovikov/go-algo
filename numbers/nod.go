package numbers

func FindNod(a int, b int) int {
	if a > b {
		return findNod(a, b)
	}

	return findNod(b, a)
}

func findNod(a int, b int) int {
	if b == 0 {
		return a
	}

	rest := a % b
	a = b
	b = rest

	return findNod(a, b)
}
