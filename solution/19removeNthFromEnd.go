package solution

// type ListNode struct {
//      Val int
//      Next *ListNode
// }

//双指针 O(L) O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hair := &ListNode{Next: head}
	pre, slow, fast := hair, head, head
	for fast != nil {
		if n < 1 {
			pre = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	pre.Next = slow.Next
	return hair.Next
}
