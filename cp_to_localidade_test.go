package ctt

import "testing"

func TestCPToLocalidade(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"1800", "Lisboa"},
		{"1800255", "Lisboa"},
		{"1800-255", "Lisboa"},
	}

	for _, c := range cases {
		got := CPToLocalidade(c.in)
		if got != c.want {
			t.Errorf("CPToLocalidade(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
