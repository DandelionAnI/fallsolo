package solution

//头插法 on o1
func reverseBetween(head *ListNode, left, right int) *ListNode {
	hair := &ListNode{Val: 0, Next: head}
	pre := hair
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return hair.Next
}
