/*
https://leetcode.com/problems/two-sum/
本题主要是掌握使用(B)Tree Map的数据结构达到O(n)的时间复杂度
不演示暴力遍历求解的方法
*/

package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	//m := make(map[int]int)
	for i := 0; i < len(nums); i++ {

	}
	return []int{-1}
}

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 5}, 9)
	fmt.Println(result)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("Wrong Answer")
	}
}
