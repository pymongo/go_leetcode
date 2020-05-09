package go_leetcode

import "testing"

func Add(x int32, y int32) int32 {
	return x + y
}

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
