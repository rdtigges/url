package url

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}
	if got, want := u.Scheme, "https"; got != want {
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawurl, got, want)
	}
	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	}
	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", rawurl, got, want)
	}
}

func TestURLHost(t *testing.T) {
	tests := map[string]struct {
		in       string // URL.Host field
		hostname string
		port     string
	}{
		"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
		"with empty port": {in: "foo.com", hostname: "foo.com", port: ""},
		"without port":    {in: "foo.com:", hostname: "foo.com", port: ""},
		"ip with port":    {in: "1.2.3.4:90", hostname: "1.2.3.4", port: "90"},
		"ip without port": {in: "1.2.3.4", hostname: "1.2.3.4", port: ""},
		// Add more tests in case of a need
	}
	for name, tt := range tests {
		t.Run(fmt.Sprintf("Hostname/%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Hostname(), tt.hostname; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
		t.Run(fmt.Sprintf("Port/%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}
