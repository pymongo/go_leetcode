package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	Nums1  []int
	M      int
	Nums2  []int
	N      int
	Result []int
}

var testCases = []TestCase{
	{Nums1: []int{1, 2, 3, 0, 0, 0}, M: 3, Nums2: []int{2, 5, 6}, N: 3, Result: []int{1, 2, 2, 3, 5, 6}},
	{[]int{1}, 1, []int{}, 0, []int{1}},
	{[]int{0}, 0, []int{1}, 1, []int{1}},
	{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1ReadIndex := m - 1
	nums2ReadIndex := n - 1
	nums1WriteIndex := m + n - 1
	for nums1ReadIndex >= 0 && nums2ReadIndex >= 0 {
		if nums1[nums1ReadIndex] > nums2[nums2ReadIndex] {
			nums1[nums1WriteIndex] = nums1[nums1ReadIndex]
			nums1ReadIndex--
			nums1WriteIndex--
		} else {
			nums1[nums1WriteIndex] = nums2[nums2ReadIndex]
			nums2ReadIndex--
			nums1WriteIndex--
		}
	}
	/*
		Input:    [2,0] 1 [1] 1
		Output:   [2,2]
		Expected: [1,2]
	*/
	// 如果nums1从后往前遍历完的时候，nums2还有剩余项
	// 则把nums2[0:nums2ReadIndex]赋值/复制到nums1[0:nums2ReadIndex]中
	// add missing elements from nums2
	if nums2ReadIndex > -1 {
		for i := 0; i <= nums2ReadIndex; i++ {
			nums1[i] = nums2[i]
		}
	}
}

func TestMerge(t *testing.T) {
	for i := 0; i < len(testCases); i++ {
		merge(testCases[i].Nums1, testCases[i].M, testCases[i].Nums2, testCases[i].N)
		if !reflect.DeepEqual(testCases[i].Nums1, testCases[i].Result) {
			fmt.Printf("Expected: %#v\n", testCases[i].Result)
			fmt.Printf("Actual  : %#v\n", testCases[i].Nums1)
			t.Errorf("Wrong Answer")
		}
	}
}

/* Best
func merge(nums1 []int, m int, nums2 []int, n int)  {

    for m > 0 && n > 0 {
        if nums1[m - 1] > nums2[n - 1] {
            nums1[m+n-1] = nums1[m-1]
            m--
        } else {
            nums1[m+n-1] = nums2[n-1]
            n--
        }
    }
    if m == 0 {
        for i := 0; i < n; i++ {
            nums1[i] = nums2[i]
        }
    }
}
*/
