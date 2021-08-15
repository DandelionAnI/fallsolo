package main

// import (
// 	"fmt"
// 	"os"
// 	//. "solution/solution"
// )

// func main() {
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

import (
	"fmt"
	"sort"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Println("true")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func formatList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	p1, cur := head.Next, head
	pre := &ListNode{Val: 0, Next: cur}
	for p1 != nil && p1.Next != nil {
		cur = cur.Next
		p1 = p1.Next
		cur.Next = p1.Next

		p1.Next = pre.Next
		pre.Next = p1
		p1 = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
		cur = cur.Next
	}
	return pre.Next
}
func ans(array []int, k int) int64 {
	n := len(array)
	var cnt int64
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if array[i]+array[j] <= k {
				cnt++
			}
		}
	}
	return cnt
}

func ans1(array []int, k int) int64 {
	n := len(array)
	var cnt int64
	sort.Ints(array)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if array[i]+array[j] <= k {
				cnt++
			} else {
				break
			}
		}
	}
	return cnt
}
