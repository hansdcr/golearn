package main

import (
	"fmt"
	"golearn/ArrayList"

)

func main1() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}


func main2() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	fmt.Println(list)
}

func main3() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for i :=0; i< 15; i++ {
		err :=list.Insert(0, "x5")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(list)
	}


}


func main4 () {
	a :=[]int{0,1,2,3,4,5}
	// var b []int
	b :=make([]int, 7)
	copy(b, a)
	fmt.Println(b)

	for i :=len(a); i>2; i-- {
		b[i] = b[i-1]
	}
	b[3]=10
	fmt.Println(b)
}

func main() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("a3")
	list.Append("a4")
	list.Append("a4")

	fmt.Println(list)
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}