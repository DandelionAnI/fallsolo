package solution

//递归 on on
func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	newhead := reverse25(head, p)
	head.Next = reverseKGroup(p,k)
	return newhead
}

func reverse25(p1, p2 *ListNode) *ListNode {
	pre := p2
	for p1 != p2 {
		next := p1.Next
		p1.Next = pre
		pre = p1
		p1 = next
	}
	return pre
}