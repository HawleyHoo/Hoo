/*
@Time : 2019-07-10 10:09
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("path:[%s]\n", pathname+"/"+fi.Name())
			//GetAllFile(pathname + "/" + fi.Name() )
		} else {
			fmt.Println(fi.Name(), fi.Sys())
			fmt.Println(fi.Mode(), fi.ModTime(), fi.Size())
		}
	}
	return err
}

func main() {
	readlinkfile()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(pwd)
	return
	pwd = "/Users/lx/Desktop/goworks/src/exchange_services/services_admin/log"
	GetAllFile(pwd)
}

func readlinkfile() {
	filename := "/Users/lx/Desktop/goworks/src/exchange_services/services_admin/log/adminweb"
	//filename := "/Users/lx/Desktop/goworks/src/exchange_services/services_admin/log/adminweb.20190702.log"
	//GetAllFile("/Users/lx/Desktop/goworks/src/exchange_services/services_admin/log")

	fileInfo, _ := lstat(filename)
	stat(filename)
	fmt.Println("mode:", fileInfo.Mode())
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println(filename, "is a symbolic link")
		readlink(filename)
		//realPath, err := filepath.EvalSymlinks("/Users/lx/Desktop/goworks/src/exchange_services/services_admin/log/adminweb")
		//if err == nil {
		//	fmt.Println("Path:", realPath)
		//}
		//fmt.Println(err)
	}
}

func readlink(filename string) {
	path, err := os.Readlink(filename)
	if err != nil {
		fmt.Println("readlink err:", err)
	}
	fmt.Println("readlink path:", path)
}

func lstat(filename string) (os.FileInfo, error) {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println("lstat err:", err)
		return fileInfo, err
	}
	fmt.Printf("Link stat info: %+v \n", fileInfo.Mode())
	return fileInfo, err
}

func stat(filename string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("stat err:", err)
		return fileInfo, err
	}
	fmt.Printf("stat info: %+v \n", fileInfo.Mode())
	return fileInfo, err
}
