package ArrayList

type StackArray interface {
	Push(interface{}) // 压入
	Pop() interface{} // 弹出
	Size() int        // 栈的大小
	Clear()           // 清空
	IsFull() bool     // 是否满了
	IsEmpty() bool    // 是否为空
}

type Stack struct {
	myarray *ArrayList
	capsieze    int // 最大范围
}

func NewArrayStack() *Stack {
	stack := new(Stack)

	stack.myarray = NewArrayList()
	stack.capsieze = 10

	return stack
}

func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.myarray.Append(data)
	}
}

func (stack *Stack) Pop() (data interface{}) {
	if !stack.IsEmpty() {
		last := stack.myarray.dataSource[stack.myarray.TheSize - 1]
		stack.myarray.Delete(stack.myarray.TheSize -1 )
		return last
	}
	return nil
}

func (stack *Stack) Size() (size int) {
	return stack.myarray.TheSize
}

func (stack *Stack) Clear() {
	stack.myarray.Clear()
}
func (stack *Stack) IsFull() bool {
	if stack.myarray.Size() >= stack.capsieze {
		return true
	} else {
		return false
	}
}
func (stack *Stack) IsEmpty() bool {
	if stack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}

}

