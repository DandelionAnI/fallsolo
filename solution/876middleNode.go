package solution
//快慢指针 ono1
func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	len := 0
	cur := head
	for cur != nil {
		len++
		cur = cur.Next
	}
	if len%2 == 0 {
		return p1.Next
	}
	return p1
}