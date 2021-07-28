package solution

//先快慢指针找到链表中点，逆序后半部分，再合并链表
//ON o1
func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	p1, p2 := head, head
	for p1.Next != nil && p1.Next.Next != nil {
		p2 = p2.Next
		p1 = p1.Next.Next
	}
	//将p1置为开头，并从中间切断，同时让p2指向反转后的头
	mid := p2.Next
	p2.Next = nil
	p2 = reve(mid)
	p1 = head
	merg(p1, p2)
}

func reve(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func merg(p1, p2 *ListNode) {
	var l1, l2 *ListNode
	for p1 != nil && p2 != nil {
		l1 = p1.Next
		p1.Next = p2
		p1 = l1

		l2 = p2.Next
		p2.Next = p1
		p2 = l2
	}
}
