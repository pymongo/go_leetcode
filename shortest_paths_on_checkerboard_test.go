package go_leetcode

import (
	"fmt"
	"testing"
)

func UniquePaths(m int, n int) int {
	// 端点是4x4，但是棋盘的格子就3x3
	max, min := m-1, n-1
	if min > max {
		max, min = min, max
	}
	result := 1
	sum := max+min
	for i:=0; i<min; i++ {
		result *= sum-i
		// 先乘后除，避免result溢出
		result /= i+1
	}
	return result
}

func TestUniquePaths(t *testing.T) {
	result := UniquePaths(23, 12)
	fmt.Println(result)
	if result != 193536720 {
		t.Errorf("uniquePaths(23, 12) = %d; want 193536720", result)
	}
}
