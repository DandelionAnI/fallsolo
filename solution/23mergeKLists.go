package solution

//
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n < 1 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	mid := n / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val > l2.Val {
// 		l2.Next = mergeTwoLists(l1, l2.Next)
// 		return l2
// 	} else {
// 		l1.Next = mergeTwoLists(l1.Next, l2)
// 		return l1
// 	}
// }
