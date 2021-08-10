package solution

type MinStack struct {
	s []Node
}
type Node struct {
	val int //data
	min int //当前最小值
}

func MSConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	val := Node{val: x, min: x}
	if len(this.s) > 0 && this.s[len(this.s)-1].min < x {
		val.min = this.s[len(this.s)-1].min
	}
	this.s = append(this.s, val)
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].val
}

func (this *MinStack) GetMin() int {
	return this.s[len(this.s)-1].min
}
