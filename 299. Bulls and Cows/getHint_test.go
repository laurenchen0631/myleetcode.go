package bullsandcows_test

import (
	"testing"

	. "./"
)

func TestGetHint(t *testing.T) {
	cases := []struct {
		in  [2]string
		out string
	}{
		{[2]string{"1807", "7810"}, "1A3B"},
		{[2]string{"1123", "0111"}, "1A1B"},
	}

	for _, c := range cases {
		res := GetHint(c.in[0], c.in[1])
		if res != c.out {
			t.Errorf("GetHint(%s, %s) == %q, want %q", c.in[0], c.in[1], res, c.out)
		}
	}
}
