package Queue

type MyQueue interface {
	Size() int                // 大小
	Front() interface{}       // 返回第一个元素
	End() interface{}         //返回最后一个元素
	IsEmpty() bool            // 是否为空
	EnQueue(data interface{}) // 入队
	DeQueue() interface{}     // 出队
	Clear()                   // 清空
}

type Queue struct {
	dataStore []interface{} //队列的数据存储
	theSize   int           // 队列的大小
}


func NewQueue() * Queue {
	myqueue :=new(Queue)
	myqueue.dataStore = make([]interface{}, 0)
	myqueue.theSize = 0
	return myqueue
}

func (myqueue Queue) Size() int  {
	return myqueue.theSize
}

func (myqueue *Queue) Front() interface{} {
	if myqueue.Size() == 0 {
		return nil
	}
	return myqueue.dataStore[0]
}

func (myqueue *Queue) End() interface{} {
	if myqueue.Size() == 0 {
		return nil
	}
	return myqueue.dataStore[myqueue.Size()-1]
}

func (myqueue *Queue) IsEmpty() bool {
	return  myqueue.theSize == 0
}

func (myqueue *Queue) EnQueue(data interface{}) {
	myqueue.dataStore = append(myqueue.dataStore, data)
	myqueue.theSize++
}

func (myqueue *Queue) DeQueue() interface{}  {
	if myqueue.Size() == 0 {
		return nil
	}

	data :=myqueue.dataStore[0]
	if myqueue.theSize > 1 {
		myqueue.dataStore = myqueue.dataStore[1:myqueue.theSize]
	} else {
		myqueue.dataStore = make([]interface{}, 0)
	}
	myqueue.theSize--
	return data
}

func (myqueue *Queue) Clear() {
	myqueue.dataStore = make([]interface{}, 0)
	myqueue.theSize = 0
}

