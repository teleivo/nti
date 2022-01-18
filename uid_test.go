package tracker

import (
	"regexp"
	"testing"
)

func TestUID(t *testing.T) {
	re := regexp.MustCompile("^[a-zA-Z][0-9a-zA-Z]{10}")

	for i := 0; i < 10; i++ {
		got, err := NewUID()
		if err != nil {
			t.Fatalf("got error %s, expected none", err)
		}
		if !re.MatchString(string(got)) {
			t.Errorf("got UID %q which does not match the expected pattern", got)
		}
	}
}
