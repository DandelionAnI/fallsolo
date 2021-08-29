package solution

type MaxStack struct {
	s []Node1
}
type Node1 struct {
	val int //data
	max int //当前最大值
}

func Constructor716() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(x int) {
	val := Node1{val: x, max: x}
	if len(this.s) > 0 && this.s[len(this.s)-1].max < x {
		val.max = this.s[len(this.s)-1].max
	}
	this.s = append(this.s, val)
}

func (this *MaxStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MaxStack) Top() int {
	return this.s[len(this.s)-1].val
}

func (this *MaxStack) peekMax() int {
	return this.s[len(this.s)-1].max
}

func (this *MaxStack) popMax() int {
	max := this.s[len(this.s)-1].max
	this.s = this.s[:len(this.s)-1]
	return max
}

// 链表实现
// class MaxStack {
//     private Node head;

//     public void push(int x) {
//         if(head == null)
//             head = new Node(x, x);
//         else
//             head = new Node(x, Math.max(x, head.max), head);
//     }

//     public void pop() {
//         head = head.next;
//     }

//     public int top() {
//         return head.val;
//     }

//     public int getMax() {
//         return head.max;
//     }

//     private class Node {
//         int val;
//         int max;
//         Node next;

//         private Node(int val, int max) {
//             this(val, max, null);
//         }

//         private Node(int val, int max, Node next) {
//             this.val = val;
//             this.max = max;
//             this.next = next;
//         }
//     }
// }
