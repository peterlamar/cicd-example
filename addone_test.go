package mathhelp

import "testing"

func TestAddone(t *testing.T) {
	for _, c := range []struct {
		in, want int
	}{
		{1, 2},
	} {
		got := addone(c.in)
		if got != c.want {
			t.Errorf("addone(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
