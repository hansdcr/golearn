package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                  // 数组的大小
	Get(index int) (interface{}, error)         // 通过索引获取数组元素
	Set(index int, newval interface{}) error    // 设置值
	Insert(index int, newval interface{}) error // 插入值
	Append(newval interface{})                  // 向数组追加一个值
	Delete(index int) error                     // 删除一个值
	Clear()                                     // 清空数组
	String() string                             // 打印数组
}

type ArrayList struct {
	dataSource [] interface{} // 存储范型数据
	TheSize    int            // 数组的大小
}

func NewArrayList() *ArrayList {
	/*
		知识点： new和make的区别
		1、都是分配内存地址，但是new返回的是指针, make返回的是类型本身
		2、范围不同，make仅用于slice,chan,map这三种类型
	*/
	list := new(ArrayList)                       // 初始化结构体
	list.dataSource = make([]interface{}, 0, 10) // 开辟空间
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < list.TheSize || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}

	return list.dataSource[index], nil
}

func (list *ArrayList) Append(newval interface{}) {
	list.dataSource = append(list.dataSource, newval)
	list.TheSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataSource)
}

func (list *ArrayList) Clear() {
	list.dataSource = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.dataSource[index] = newval

	return nil
}

func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.dataSource = append(list.dataSource[:index], list.dataSource[index+1:]...)
	list.TheSize--

	return nil
}

func (list *ArrayList) checkFull() {
	if list.TheSize == cap(list.dataSource) { // 数据大小和实际使用容量相等
		newDataSource := make([]interface{}, 2*list.TheSize) //开辟两倍的地址空间
		copy(newDataSource, list.dataSource)
		list.dataSource = newDataSource
	}
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.checkFull() //检测内存，如果满了，自动追加

	list.dataSource = list.dataSource[:list.TheSize+1] // 插入数据，内存移动一位
	for i := list.TheSize; i > index; i-- { //从后往前移动
		list.dataSource[i] = list.dataSource[i-1]
	}
	list.dataSource[index] = newval // 插入数据
	list.TheSize++                  // 索引追加

	return nil
}
