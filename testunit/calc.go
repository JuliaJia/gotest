package testunit

import "testing"

func TestAdd(add func(a int, b int) int, t *testing.T) {
	if 3 == add(1, 2) {
		t.Error("1 + 2 != 3")
	}
}
