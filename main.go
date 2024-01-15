package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// 定义当前目录信息
type currentDirInfo struct {
	currentDirs  []string
	currentFiles []string
}

func getCurrentDir(path string) (currentDirInfo, error) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
	   return currentDirInfo{}, nil
	}
	var currentDirs []string
	var currentFiles []string
	for _, info := range infos {
	   fmt.Printf("Name:%-30s    Size:%-10d字节    Modtime:%s    IsDir:%v\n", info.Name(), info.Size(), info.ModTime().Format("2006-01-02 03:04:05"), info.IsDir())
	   abs, err := filepath.Abs(info.Name())
	   if err != nil {
		  return currentDirInfo{}, nil
	   }
	   if info.IsDir() {
		  currentDirs = append(currentDirs, abs)
	   } else {
		  currentFiles = append(currentFiles, abs)
	   }
	}
	fmt.Printf("=======================当前文件夹==============================\n")
	for _, v := range currentDirs {
	   fmt.Printf("当前文件夹:%s\n", v)
	}
	fmt.Printf("========================当前文件=============================\n")
	for _, v := range currentFiles {
	   fmt.Printf("当前文件:%s\n", v)
	}
 
	return currentDirInfo{currentDirs: currentDirs, currentFiles: currentFiles}, nil
 }
 func main() {
	_, err := getCurrentDir(".")
	if err != nil {
	   return
	}
 }