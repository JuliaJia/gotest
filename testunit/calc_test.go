package testunit

import "testing"

func Add(a int, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	if 3 != Add(1, 2) {
		t.Error("1 + 2 != 3")
	}
}
