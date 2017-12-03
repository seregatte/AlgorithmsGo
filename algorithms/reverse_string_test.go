package algorithms

import "testing"

func TestReverseString(t *testing.T) {
	v := ReverseString("abcd")
	if v != "dcba" {
		t.Error("Expected dcba, got ", v)
	}
}
