package ArrayList

type StackArrayX interface {
	Push(interface{}) // 压入
	Pop() interface{} // 弹出
	Size() int        // 栈的大小
	Clear()           // 清空
	IsFull() bool     // 是否满了
	IsEmpty() bool    // 是否为空
}

type StackX struct {
	myarray *ArrayList
	// capsieze    int // 最大范围
	Myit Iterator
}

func NewArrayListStackX() *StackX {
	stack := new(StackX)

	stack.myarray = NewArrayList() //数组
	stack.Myit = stack.myarray.Iterator() //迭代

	return stack
}

func (stack *StackX) Push(data interface{}) {
	if !stack.IsFull() {
		stack.myarray.Append(data)
	}
}

func (stack *StackX) Pop() (data interface{}) {
	if !stack.IsEmpty() {
		last := stack.myarray.dataSource[stack.myarray.TheSize - 1]
		stack.myarray.Delete(stack.myarray.TheSize -1 )
		return last
	}
	return nil
}

func (stack *StackX) Size() (size int) {
	return stack.myarray.TheSize
}

func (stack *StackX) Clear() {
	stack.myarray.Clear()
	stack.myarray.TheSize = 0
}
func (stack *StackX) IsFull() bool {
	if stack.myarray.Size() >= 10 {
		return true
	} else {
		return false
	}
}
func (stack *StackX) IsEmpty() bool {
	if stack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}

}

