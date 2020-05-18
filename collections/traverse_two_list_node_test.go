package collections

import (
	"reflect"
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
func listNodeToSlice(listNode *ListNode) []int {
	var numbers []int
	for listNode != nil {
		numbers = append(numbers, listNode.Val)
		listNode = listNode.Next
	}
	return numbers
}

func TestHelperMethods(t *testing.T) {
	result := sliceToListNode([]int{1, 2, 3})
	if !reflect.DeepEqual(listNodeToSlice(result), []int{1, 2, 3}) {
		t.Errorf("Wrong Answer")
	}
}
