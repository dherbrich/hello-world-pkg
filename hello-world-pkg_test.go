package hello

import (
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		in, want string
	}{
		{"denny", "Hello denny!"},
		{"nadine", "Hello nadine!"},
		{"!&/%$", "Hello !&/%$!"},
	}

	for _, test := range tests {
		if out := SayHello(test.in); out != test.want {
			t.Errorf("SayHello(%q): got: %q <> want: %q", test.in, out, test.want)
		}
	}
}
