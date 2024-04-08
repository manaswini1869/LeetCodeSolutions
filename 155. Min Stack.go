type MinStack struct {
    vals []int
    min_val int    
}


func Constructor() MinStack {
    return MinStack{
        vals: []int{},
        min_val: 1<<63-1,
    }
    
}


func (this *MinStack) Push(val int)  {
    if this.min_val > val {
        this.min_val = val
    }
    this.vals = append(this.vals, val)
    
}


func (this *MinStack) Pop()  {
    temp := this.vals[len(this.vals)-1]
    this.vals = this.vals[:len(this.vals)-1]
    if this.min_val == temp {
        this.min_val = 1<<63-1
        for _,i := range(this.vals) {
            if i < this.min_val {
                this.min_val = i
            }
        }
    }

}


func (this *MinStack) Top() int {
    return this.vals[len(this.vals)-1]
    
}


func (this *MinStack) GetMin() int {
    return this.min_val    
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */