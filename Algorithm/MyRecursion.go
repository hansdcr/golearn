package Algorithm

import "golearn/ArrayList"

func MyRecursion(num int) int {
	mystack :=ArrayList.NewArrayListStackX()
	mystack.Push(num)
	last :=0

	for !mystack.IsEmpty() {
		data := mystack.Pop()
		if data == 0 {
			last +=0
		} else {
			last +=data.(int)
			mystack.Push((data.(int)-1))
		}
	}

	return last
}
