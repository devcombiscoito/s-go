package app

import "testing"

func TestIsNewer(t *testing.T) {
	cases := []struct {
		remote  string
		current string
		want    bool
	}{
		{"1.0.1", "1.0.0", true},
		{"1.1.0", "1.0.9", true},
		{"2.0.0", "1.9.9", true},
		{"1.0.0", "1.0.0", false},
		{"1.0.0", "1.0.1", false},
		{"v1.2.3", "1.2.3", false},
		{"v1.2.4", "1.2.3", true},
		{"1.2", "1.2.0", false},
		{"1", "1.0.1", false},
	}

	for _, c := range cases {
		if got := isNewer(c.remote, c.current); got != c.want {
			t.Fatalf("isNewer(%q, %q) = %v, want %v", c.remote, c.current, got, c.want)
		}
	}
}
