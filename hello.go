package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	pre := &ListNode{Val: 1}
	// fmt.Scan(&pre.Val)
	// head := pre
	// for {
	// 	node := &ListNode{}
	// 	head.Next = node
	// 	_, err := fmt.Scan(&node.Val)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	head = head.Next

	// }
	head := &ListNode{}
	head = pre
	// pre.Next = head
	for i := 2; i < 6; i++ {
		node := &ListNode{Val: i}
		head.Next = node
		head = head.Next
	}
	n := 0
	fmt.Scanln(&n)
	pre = removeNthFromEnd(pre, n)
	for pre != nil {
		fmt.Print(pre.Val)
		pre = pre.Next
	}
}

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

// func Code() {
// }

// func createfile() {
// 	num := os.Args[1]
// 	topic := os.Args[2]

// 	f, err := os.Create("./solution/" + num + topic + ".go")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		f.Write([]byte("package solution\n"))
// 		f.Write([]byte("\n"))
// 		f.Write([]byte("func " + topic + "() {\n"))
// 		f.Write([]byte("}\n"))
// 	}
// 	defer f.Close()
// }
