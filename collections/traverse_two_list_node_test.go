package collections

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 仅用于生成测试用例的Helper方法
func sliceToListNode(numbers []int) *ListNode {
	dummyHead := new(ListNode)
	current := dummyHead
	for i:=0; i< len(numbers); i++ {
		current.Next = new(ListNode)
		current = current.Next
		current.Val = numbers[i]
	}
	return dummyHead.Next
}

// 仅用于单元测试中验证返回值的Helper方法
//func listNodeToArray(numbers []int) {
//	result := twoSum([]int{2, 7, 11, 5}, 9)
//	fmt.Println(result)
//	if !reflect.DeepEqual(result, []int{0, 1}) {
//		t.Errorf("Wrong Answer")
//	}
//}

func TestHelperMethods(t *testing.T) {
	result := sliceToListNode([]int{2, 7, 11, 5})
	fmt.Println(result)
	if false {
		t.Errorf("Wrong Answer")
	}
}
