package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	i, j := x, 0
	for i > 0 {
		j = j*10 + i%10
		i = i / 10
	}
	return x == j
}
