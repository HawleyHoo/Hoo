package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	/*图片：
	<img...src="(http[\s\S]+?)"...>
	<img...alt="中国军团“备战”维密 大秀前先来一波颜值霸屏！"...src="http://cms-bucket.nosdn.127.net/2018/11/09/0e08464f5c9f4448ad5db7683283f571.jpeg?imageView&amp;thumbnail=200y125&amp;quality=85">
	*/
	//reImg = `<img[\s\S]+?src="(http[\s\S]+?)"[\s\S]*?>`
	//reImg = `<img[\s\S]+?src="(http[\s\S]+?)"`
	reImg    = `<img.+?src="(http.+?)".*?>`
	reImgAlt = `<img.+?alt="(.+?)"`

	/*img标签中的alt属性*/
	reAlt = `alt="([\s\S]+?)"`

	/*
		图片链接中的图片名称
		http://cms-bucket.nosdn.127.net/2018/11/09/7e88b8526ff141129809d8ae7c718e51.jpeg?imageView&thumbnail=185y116&quality=85
		http://img2.money.126.net/chart/hs/time/180x120/0000001.png
		http://cms-bucket.nosdn.127.net/2018/05/31/bc7d30ff42194c35a4743834a77ec97b.png?imageView&thumbnail=90y90&quality=85
	*/
	reImgName = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(webp)|(swf)|(ico)))`
)

func GetHtml(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}

/*爬取图片链接*/
func spiderImg() {
	html := GetHtml("http://www.163.com")

	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))
	for _, ret := range rets {
		fmt.Println(ret[1])
	}
}

/*爬取图片链接中的alt属性*/
func spiderImgAlt() {
	html := GetHtml("http://www.163.com")

	//将gbk转utf8
	bytes := ConvertToByte(html, "gbk", "utf8")
	html = string(bytes)
	//fmt.Println(html)

	re := regexp.MustCompile(reImgAlt)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))
	for _, ret := range rets {
		fmt.Println(ret[1])
	}
}

/*爬取图片链接及其alt*/
func spiderImgNameAlt() {
	html := GetHtml("http://www.163.com")

	//将gbk转utf8
	bytes := ConvertToByte(html, "gbk", "utf8")
	html = string(bytes)
	//fmt.Println(html)

	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))
	for i, ret := range rets {
		imgTag := ret[0]
		fmt.Println(i, imgTag, "\n", GetImgNameFromTag(imgTag, ret[1], imgDir), "\n")
	}
}

/*爬取图片链接中的文件名*/
func spiderImgName() {
	html := GetHtml("http://www.163.com")

	//将gbk转utf8
	bytes := ConvertToByte(html, "gbk", "utf8")
	html = string(bytes)
	//fmt.Println(html)

	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))
	for _, ret := range rets {
		//imgUrl := ret[1]
		//fmt.Println(ret[0],imgUrl,GetImgNameFromImgurl(imgUrl))
		fmt.Println(ret[0])
	}
}

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func main011() {
	spiderImgName()
}
