package Algorithm

import (
	"errors"
	"io/ioutil"
)

func GetAll(path string, files []string, level int) ([]string, error){
	levelstr :=""
	if level == 1 {
		levelstr="--"
	}else{
		for ;level>1; level-- {
			levelstr +="|--"
		}
	}

	read, err :=ioutil.ReadDir(path)
	if err !=nil {
		return files, errors.New("文件夹不可读取")
	}

	for _,fi :=range read {
		if fi.IsDir(){ //如果是文件夹
			fulldir := path+"/"+fi.Name()
			files = append(files, levelstr+fulldir)
			files, _ = GetAll(fulldir, files, level+1)
		} else {
			fulldir := path+"/"+fi.Name()
			files = append(files, levelstr+fulldir)
		}
	}

	return files, nil
}