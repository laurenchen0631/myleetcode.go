package longestpalindrome_test

import (
	"testing"

	. "./"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"abccccdd", 7},
		{"ccc", 3},
	}

	for _, c := range cases {
		res := LongestPalindrome(c.in)
		if res != c.out {
			t.Errorf("LongestPalindrome(%q) == %q, want %q", c.in, res, c.out)
		}
	}
}
