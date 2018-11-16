package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	QDMStockOrder struct {
		StoreCode string `json:"StoreCode"`
		PKUCode   string `json:"PKUCode"`
		Qty       int    `json:"Qty"`
	}

	QDMFile struct {
		Count int             `json:"count"`
		List  []QDMStockOrder `json:"list"`
	}
)

//func testMarshal() []byte {
//	user := User{}
//	data, err := json.Marshal(user)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return data
//}

func testRead() []byte {
	fp, err := os.OpenFile("./ordercontent.json", os.O_RDONLY, 0766)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	//data := make([]byte, 100)
	n, err := ioutil.ReadAll(fp)
	//n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(n))
	return n
}

func readFile() []byte {
	b, err := ioutil.ReadFile("./h_base/ordercontent.json")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(b)
	str := string(b)
	fmt.Println(str)
	return b
}

func testWrite(data []byte) {
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//var data []byte
	//data1 := testRead()
	data1 := readFile()
	data1 = bytes.TrimPrefix(data1, []byte("\xef\xbb\xbf"))
	res := QDMFile{}
	err := json.Unmarshal(data1, &res)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Printf("file: %+v\n", res)
	//l := len(res.List)
	//for i := 0; i < l; i++  {
	//
	//}
	list := []QDMStockOrder{}
	for _, v := range res.List {
		if len(v.PKUCode) == 0 || len(v.StoreCode) == 0 || v.Qty == 0 {
			fmt.Println("错误数据：", v)
			continue
		}
		list = append(list, v)
	}
	fmt.Println("res :", list, len(list))
	//err = model.StockAddOrder(res.List)
	//fmt.Println("db err", err)

	//fmt.Println(string(data))
	//testWrite(data)
	//data = testRead()
	//testUnmarshal(data)
}

func Ioutil(name string) {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use ioutil.ReadFile to read a file:", result)
	}
}

func OsIoutil(name string) {
	if fileObj, err := os.Open(name); err == nil {
		//if fileObj,err := os.OpenFile(name,os.O_RDONLY,0644); err == nil {
		defer fileObj.Close()
		if contents, err := ioutil.ReadAll(fileObj); err == nil {
			result := strings.Replace(string(contents), "\n", "", 1)
			fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file :", result)
		}

	}
}

func FileRead(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
		buf := make([]byte, 1024)
		if n, err := fileObj.Read(buf); err == nil {
			fmt.Println("The number of bytes read:"+strconv.Itoa(n), "Buf length:"+strconv.Itoa(len(buf)))
			result := strings.Replace(string(buf), "\n", "", 1)
			fmt.Println("Use os.Open and File's Read method to read a file:", result)
		}
	}
}

func BufioRead(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
		reader := bufio.NewReader(fileObj)
		//使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
		if result, err := reader.ReadString(byte('@')); err == nil {
			fmt.Println("使用ReadSlince相关方法读取内容:", result)
		}
		//注意:上述ReadString已经将buffer中的数据读取出来了，下面将不会输出内容
		//需要注意的是，因为是将文件内容读取到[]byte中，因此需要对大小进行一定的把控
		buf := make([]byte, 1024)
		//读取Reader对象中的内容到[]byte类型的buf中
		if n, err := reader.Read(buf); err == nil {
			fmt.Println("The number of bytes read:" + strconv.Itoa(n))
			//这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
			fmt.Println("Use bufio.NewReader and os.Open read file contents to a []byte:", string(buf))
		}

	}
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}

//使用os.OpenFile()相关函数打开文件对象，并使用文件对象的相关方法进行文件写入操作
//清空一次文件
func WriteWithFileWrite(name, content string) {
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.", content)
	}
	contents := []byte(content)
	if _, err := fileObj.Write(contents); err == nil {
		fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.", content)
	}
}

//使用io.WriteString()函数进行数据的写入
func WriteWithIo(name, content string) {
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.", content)
	}
}

//使用bufio包中Writer对象的相关方法进行数据的写入
func WriteWithBufio(name, content string) {
	if fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriterSize(fileObj, 4096)
		//
		if _, err := writeObj.WriteString(content); err == nil {
			fmt.Println("Successful appending buffer and flush to file with bufio's Writer obj WriteString method", content)
		}

		//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
		buf := []byte(content)
		if _, err := writeObj.Write(buf); err == nil {
			fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.", content)
			if err := writeObj.Flush(); err != nil {
				panic(err)
			}
			fmt.Println("Successful flush the buffer data to file ", content)
		}
	}
}

/*
作者：BGbiao
链接：https://www.jianshu.com/p/7790ca1bc8f6
來源：简书
简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
*/
