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
	for i := 0; i < len(numbers); i++ {
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	current := dummyHead
	sumOrCarry := 0
	for {
		if l1 != nil && l2 != nil {
			current.Next = new(ListNode)
			current = current.Next
			sumOrCarry = sumOrCarry + l1.Val + l2.Val
			current.Val = sumOrCarry % 10
			sumOrCarry /= 10
			l1, l2 = l1.Next, l2.Next
		} else if l1 != nil && l2 == nil {
			current.Next = new(ListNode)
			current = current.Next
			sumOrCarry = sumOrCarry + l1.Val
			current.Val = sumOrCarry % 10
			sumOrCarry /= 10
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			current.Next = new(ListNode)
			current = current.Next
			sumOrCarry = sumOrCarry + l2.Val
			current.Val = sumOrCarry % 10
			sumOrCarry /= 10
			l2 = l2.Next
		} else {
			break
		}
	}
	// 如果还有进位，应对testCase: [5],[5] => [0, 1]
	if sumOrCarry > 0 {
		current.Next = new(ListNode)
		current = current.Next
		current.Val = sumOrCarry
	}
	return dummyHead.Next
}

func TestAddTwoNumbersTestCase2(t *testing.T) {
	ln1 := sliceToListNode([]int{5})
	ln2 := sliceToListNode([]int{5})
	result := addTwoNumbers(ln1, ln2)
	if !reflect.DeepEqual(listNodeToSlice(result), []int{0, 1}) {
		t.Errorf("Wrong Answer")
	}
}

func TestAddTwoNumbersTestCase1(t *testing.T) {
	ln1 := sliceToListNode([]int{2, 4, 3})
	ln2 := sliceToListNode([]int{5, 6, 4})
	result := addTwoNumbers(ln1, ln2)
	if !reflect.DeepEqual(listNodeToSlice(result), []int{7, 0, 8}) {
		t.Errorf("Wrong Answer")
	}
}
