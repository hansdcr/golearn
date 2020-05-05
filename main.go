package main

import (
	"fmt"
	"test/ArrayList"
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

func main() {
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
