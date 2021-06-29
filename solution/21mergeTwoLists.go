package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//递归 O(n+m) o(n+m)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

//迭代 O(n+m) o(1)
// func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head := &ListNode{Val: 0}
// 	hair := head
// 	for l1 != nil && l2 != nil {
// 		if l1.Val > l2.Val {
// 			head.Next = l2
// 			head = head.Next
// 			l2 = l2.Next
// 		} else {
// 			head.Next = l1
// 			head = head.Next
// 			l1 = l1.Next
// 		}
// 	}
// 	if l1 == nil {
// 		head.Next = l2
// 	} else {
// 		head.Next = l1
// 	}
// 	return hair.Next
// }
