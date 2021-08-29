package solution

type MyStack struct {
	q1 []int
	q2 []int
}

//构造函数
func Constructor225() MyStack {
	return MyStack{[]int{}, []int{}}
}

//元素入栈
func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

//元素出栈
func (this *MyStack) Pop() int {
	//将队列1中除最后一个元素的所有元素加入队列2中，并从头结点删除
	//必须定义n，为队列1初始长度，后续会改变
	n := len(this.q1)
	for i := 0; i < n-1; i++ {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}
	top := this.q1[0]
	//重置队列1和队列2，有元素的始终是队列1
	this.q1 = this.q2
	this.q2 = nil

	return top
}

//获取栈顶元素
func (this *MyStack) Top() int {
	//将队列头元素弹出，再放进去
	top := this.Pop()
	this.q1 = append(this.q1, top)
	return top
}

//判断栈是否空
func (this *MyStack) Empty() bool {
	if len(this.q1) == 0 {
		return true
	}
	return false
}
