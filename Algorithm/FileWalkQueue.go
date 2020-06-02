package Algorithm

import (
	"fmt"
	"golearn/Queue"
	"io/ioutil"
)

func GetAllX (path string) {
	q :=Queue.NewQueue()
	q.EnQueue(path)

	files := []string{}

	if !q.IsEmpty() {

		for ;;{
			getpath := q.DeQueue()
			if getpath == nil {
				break
			} else {
				read, _ :=ioutil.ReadDir(getpath.(string))
				for _, item := range read {
					if item.IsDir() {
						fulldir := getpath.(string) + "/" + item.Name()
						q.EnQueue(fulldir)
					} else {
						fulldir := getpath.(string) + "/" + item.Name()
						files = append(files, fulldir)
					}
				}
			}
		}
	}

	for i:=0;i<len(files);i++ {
		fmt.Println(files[i])
	}
}