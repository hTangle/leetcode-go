package __add_two_number

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNodeFromNumber(number int) *ListNode {
	if number == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	val := &ListNode{
		Val: number % 10,
	}
	header := val
	number = number / 10
	for number > 0 {
		val.Next = &ListNode{
			Val: number % 10,
		}
		val = val.Next
		number = number / 10
	}
	return header
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := 0
	header := &ListNode{0, nil}
	list, list1, list2 := header, l1, l2
	for list1 != nil || list2 != nil {
		if list1 != nil {
			extra += list1.Val
			list1 = list1.Next
		}
		if list2 != nil {
			extra += list2.Val
			list2 = list2.Next
		}
		list.Next = &ListNode{extra % 10, nil}
		list = list.Next
		extra = extra / 10
	}
	if extra == 1 {
		list.Next = &ListNode{1, nil}
	}
	return header.Next
}
