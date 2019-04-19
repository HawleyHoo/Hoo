package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chSem      = make(chan int, 10)
	downloadWG sync.WaitGroup
	randomMT   sync.Mutex

	//图片存储地址
	imgDir = `D:\meizi3\`
)

/*生成[start,end)之间的随机数*/
func GetRandomInt(start, end int) int {
	randomMT.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + r.Intn(end-start)
	randomMT.Unlock()
	return ret
}

/*生成时间戳_随机数文件名*/
func GetRandomName() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	randomNum := strconv.Itoa(GetRandomInt(1000, 10000))
	return timestamp + "_" + randomNum
}

/*获得页面上的全部图片链接*/
func GetPageImgurls(url string) []string {
	html := GetHtml(url)
	//fmt.Println(html)

	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))

	imgUrls := make([]string, 0)
	for _, ret := range rets {
		imgUrl := ret[1]
		imgUrls = append(imgUrls, imgUrl)
	}
	return imgUrls
}

/*获得页面上的全部图片信息（链接+文件名）*/
func GetPageImginfos(url string) []map[string]string {

	html := GetHtml(url)
	//html = string(ConvertToByte(html, "gbk", "utf8"))
	//fmt.Println(html)

	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))

	imginfos := make([]map[string]string, 0)
	for _, ret := range rets {
		imgInfo := make(map[string]string)
		imgUrl := ret[1]
		imgInfo["url"] = imgUrl
		imgInfo["filename"] = GetImgNameFromTag(ret[0], imgUrl, imgDir)

		imginfos = append(imginfos, imgInfo)
	}

	return imginfos
}

var httpClient http.Client

func init() {
	httpClient = http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {

				//设置连接请求超时时间
				conn, err := net.DialTimeout(netw, addr, time.Second*1)
				if err != nil {
					return nil, err
				}

				//设置连接的读写超时时间
				deadline := time.Now().Add(1 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
}

func DownloadImgWithClient(url string, filename string) {
	fmt.Println("DownloadImgWithClient...")
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(filename, "下载失败！")
		return
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println(filename, "下载失败！")
	}

}

func DownloadImg(url string, filename string) {
	fmt.Println("DownloadImg...")
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)

	err := ioutil.WriteFile(filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println(filename, "下载失败！")
	}

}

func DownloadImgAsync(url, filename string) {
	downloadWG.Add(1)
	go func() {
		chSem <- 123
		DownloadImg(url, filename)
		<-chSem
		downloadWG.Done()
	}()

	downloadWG.Wait()
}

/*
从<img>标签中提取文件名（含地址）：
有alt，使用alt做文件名，没有使用时间戳_随机数做文件名
参数：
imgTag 图片<img>标签
imgDir 目录位置
suffix 文件名后缀
*/
func GetImgNameFromTag(imgTag, imgUrl, imgDir string) string {
	var filename string

	//获得图片格式
	imgName := GetImgNameFromImgurl(imgUrl)
	suffix := ".jpg"
	if imgName != "" {
		suffix = imgName[strings.LastIndex(imgName, "."):]
	}

	//尝试从imgTag中提取alt
	re := regexp.MustCompile(reAlt)
	rets := re.FindAllStringSubmatch(imgTag, 1)

	if len(rets) > 0 && imgName != "" {
		//首选alt
		alt := rets[0][1]
		alt = strings.Replace(alt, ":", "_", -1)
		filename = alt + imgName
	} else if imgName != "" {
		//次选链接中的文件名
		filename = imgName
	} else {
		//最末时间戳+随机数
		filename = GetRandomName() + suffix
	}
	filename = imgDir + filename
	return filename
}

/*从imgUrl中摘取图片名称*/
func GetImgNameFromImgurl(imgUrl string) string {
	re := regexp.MustCompile(reImgName)
	rets := re.FindAllStringSubmatch(imgUrl, -1)
	if len(rets) > 0 {
		return rets[0][1]
	} else {
		return ""
	}
}

func main021() {
	//imgurls := GetPageImgurls("http://www.163.com")
	imginfos := GetPageImginfos("http://www.163.com")

	for _, imginfoMap := range imginfos {
		//DownloadImg(imgUrl)
		DownloadImgAsync(imginfoMap["url"], imginfoMap["filename"])
	}
}
