/*
https://leetcode.com/problems/two-sum/
本题主要是掌握使用(B)Tree Map的数据结构达到O(n)的时间复杂度
bitwise_补码的解法收录在java_leetcode中
不记录暴力遍历求解的方法
*/

package go_leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	// key:   target-nums[i]
	// value: index of nums
	sumTracker := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if oldIndex, keyExist := sumTracker[nums[i]]; keyExist {
			return []int{oldIndex, i}
		}
		sumTracker[target-nums[i]] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 5}, 9)
	t.Log(result)
	fmt.Println(result)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("Wrong Answer")
	}
}
