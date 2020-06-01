package Algorithm

import (
	"fmt"
	"golearn/StackArray"
	"io/ioutil"
)

func FileRecursion(path string) {
	/**
	1、给出一个第一级的路径，判断路径是否是目录，是目录压入栈中； 不是则放入文件路径数组中
	2、如果是目录，则读取目录下的所有路径，然后判断是目录还是文件， 目录的话继续递归调用，文件的话放入路径数组中
	*/
	// path :="/Users/11091752/Desktop/2020年文档整理"
	files := []string{}
	mystack := StackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		getpath := mystack.Pop().(string)
		files = append(files, getpath)
		read, _ := ioutil.ReadDir(getpath) //读取文件夹下所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name() //构造新的路径
				files = append(files, fulldir)
				mystack.Push(fulldir)
			} else {
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir)
			}
		}
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
