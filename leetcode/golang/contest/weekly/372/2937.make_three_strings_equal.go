package weekly372

func FindMinimumOperations(s1 string, s2 string, s3 string) int {
	// if the first characters of there strings are not same, return -1 directly
	if s1[0] != s2[0] || s1[0] != s3[0] || s2[0] != s3[0] {
		return -1
	}

	// count the number of same characters

	n1, n2, n3 := len(s1), len(s2), len(s3)
	i := 1
	for i < n1 && i < n2 && i < n3 && s1[i] == s2[i] && s2[i] == s3[i] {
		i++
	}

	return n1 + n2 + n3 - 3*i
}
