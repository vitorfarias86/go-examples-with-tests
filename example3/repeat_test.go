package example3

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("A")
	expected := "AAAAA"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}
