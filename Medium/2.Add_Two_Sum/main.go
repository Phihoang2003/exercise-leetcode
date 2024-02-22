package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		if temp.Val > 9 {
			temp.Val -= 10
			temp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
		}
		temp = temp.Next
	}
	return result
}
func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		println(result.Val)
		result = result.Next
	}
}
