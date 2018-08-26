package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
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

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("ccc"))

}
