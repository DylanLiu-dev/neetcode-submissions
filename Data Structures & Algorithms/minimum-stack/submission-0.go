type MinStack struct {
	vals []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		vals: []int{},
		min: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	n := len(this.min)
	if n==0 || val < this.min[n-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[n-1])
	}
}

func (this *MinStack) Pop() {
	n := len(this.vals)
	this.vals = this.vals[0:n-1]
	this.min = this.min[0:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.vals)
	res := this.vals[n-1]
	return res
}

func (this *MinStack) GetMin() int {
	n := len(this.min)
	return this.min[n-1]
}
