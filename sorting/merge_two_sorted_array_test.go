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
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	nums1ReadIdx := m - 1
	nums2ReadIdx := n - 1
	nums1WriteIdx := len(nums1) - 1
	for nums1ReadIdx >= 0 {
		if nums1[nums1ReadIdx] > nums2[nums2ReadIdx] {
			nums1[nums1WriteIdx] = nums1[nums1ReadIdx]
			nums1ReadIdx--
			nums1WriteIdx--
		} else {
			nums1[nums1WriteIdx] = nums2[nums2ReadIdx]
			if nums2ReadIdx == 0 {
				break
			}
			nums2ReadIdx--
			nums1WriteIdx--
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
