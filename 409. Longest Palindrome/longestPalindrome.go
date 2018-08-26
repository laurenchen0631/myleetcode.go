package longestpalindrome

// LongestPalindrome Given a string which consists of lowercase or uppercase letters
// find the length of the longest palindromes that can be built with those letters.
func LongestPalindrome(s string) int {
	dict := make(map[rune]int)
	for _, char := range s {
		dict[char]++
	}

	hasOddsCounter := false
	palindromeLen := 0
	for _, charCounter := range dict {
		if charCounter&1 == 1 {
			hasOddsCounter = true
		}
		palindromeLen += (charCounter / 2) * 2
	}

	if hasOddsCounter {
		palindromeLen++
	}

	return palindromeLen
}
