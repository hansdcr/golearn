package StackArray

type StackArray interface {
	Push(interface{}) // 压入
	Pop() interface{} // 弹出
	Size() int        // 栈的大小
	Clear()           // 清空
	IsFull() bool     // 是否满了
	IsEmpty() bool    // 是否为空
}

type Stack struct {
	dataSource  []interface{}
	capsieze    int // 最大范围
	currentsize int // 实际大小
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.dataSource = make([]interface{}, 0, 10)
	stack.capsieze = 10
	stack.currentsize = 0
	return stack
}

func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataSource = append(stack.dataSource, data) // 压入数据
		stack.currentsize++
	}
}

func (stack *Stack) Pop() (data interface{}) {
	if !stack.IsEmpty() {
		last := stack.dataSource[stack.currentsize-1]             // 最后一个数据
		stack.dataSource = stack.dataSource[:stack.currentsize-1] //删除最后一个数据
		stack.currentsize--
		return last
	}
	return nil
}

func (stack *Stack) Size() (size int) {
	return stack.currentsize
}

func (stack *Stack) Clear() {
	stack.dataSource = make([]interface{}, 0, 10)
	stack.capsieze = 10
	stack.currentsize = 0
}
func (stack *Stack) IsFull() bool {
	if stack.currentsize >= stack.capsieze {
		return true
	} else {
		return false
	}
}
func (stack *Stack) IsEmpty() bool {
	if stack.currentsize == 0 {
		return true
	} else {
		return false
	}

}
