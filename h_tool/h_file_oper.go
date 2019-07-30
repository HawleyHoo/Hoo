/*
@Time : 2019-07-25 21:10
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

//一些常用的文件操作函数

func main() {
	dir := "/Users/lx/Desktop/release"
	ver := "admin" + "1.0.0"
	path := path.Join(dir, ver)
	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	fmt.Println(err)
}

func main0726() {
	//创建目录
	//perm权限设置，os.ModePerm为0777
	err := os.Mkdir("./tmp", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	//创建多级目录
	err2 := os.MkdirAll("./a/b/c", os.ModePerm)
	if err2 != nil {
		log.Fatal(err2)
	}

	//等个3秒，看目录是否创建成功
	time.Sleep(time.Second * 3)

	//删除目录
	//如果目录下有文件或其他目录会出错
	err3 := os.Remove("./tmp")
	if err3 != nil {
		log.Fatal(err3)
	}

	//删除多级目录
	//如果是单个名称，则删除所有子目录
	err4 := os.RemoveAll("./a/b/c")
	if err4 != nil {
		log.Fatal(err4)
	}

	//文件操作
	//创建文件
	//Create会返回一个文件对象，默认权限0666
	file1, err5 := os.Create("./1.log")
	defer file1.Close()
	if err5 != nil {
		log.Fatal(err5)
	}

	//通过文件描述符(文件句柄)来创建文件
	file2 := os.NewFile(file1.Fd(), "./2.log")
	defer file2.Close()

	//打开文件
	//Open内部调用OpenFile，以只读方式打开
	file3, err6 := os.Open("./1.txt")
	defer file3.Close()
	if err6 != nil {
		log.Fatal(err6)
	}

	//OpenFile指定文件打开方式，只读，只写，读写和权限
	file4, err7 := os.OpenFile("./2.txt", os.O_RDWR, 0666)
	defer file4.Close()
	if err7 != nil {
		log.Fatal(file4)
	}

	//向文件写入数据
	file4.Write([]byte("我写入数据"))

	//在指定位置写入数据
	file4.WriteAt([]byte("指定位置写入数据"), 1024)

	//直接写入string
	file4.WriteString("你好世界")

	//读取文件数据

	//将文件偏移量设置为开头
	file4.Seek(0, 0)
	tmp := make([]byte, 256)
	//Read读取数据到[]byte中
	for n, _ := file4.Read(tmp); n != 0; n, _ = file4.Read(tmp) {
		fmt.Println(string(tmp))
	}

	//ReadAt从off指定位置读取数据到[]byte中
	file4.Seek(0, 0)
	var off int64 = 0
	for {
		n, _ := file4.ReadAt(tmp, off)
		off += int64(n)
		if n == 0 {
			break
		}

		fmt.Println(string(tmp))
	}

	//删除文件
	//Go中删除文件和删除文件夹同一个函数
	err8 := os.Remove("./1.txt")
	if err8 != nil {
		log.Fatal(err8)
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
