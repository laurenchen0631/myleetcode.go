package removeduplicate_test

import (
	"testing"

	. "./"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
	}

	for _, c := range cases {
		res := RemoveDuplicateLetters(c.in)
		if res != c.out {
			t.Errorf("GetRemoveDuplicateLetters(%s) == %s, want %s", c.in, res, c.out)
		}
	}
}
