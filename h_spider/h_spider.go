package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	wg4Imginfo  sync.WaitGroup
	wg4Download sync.WaitGroup
	chImgMaps   = make(chan map[string]string, 1000)
)

func main() {
	baseUrl := "https://www.duotoo.com/zt/rbmn/index"
	for i := 1; i < 5; i++ {
		var url string
		if i != 1 {
			url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		} else {
			url = baseUrl + ".html"
		}

		wg4Imginfo.Add(1)
		go func(theUrl string) {
			GetImginfosFromPage(theUrl)
			wg4Imginfo.Done()
		}(url)
	}

	go func() {
		wg4Imginfo.Wait()
		close(chImgMaps)
		//fmt.Println("chImgMaps closed!")
	}()

	for imgMap := range chImgMaps {
		//fmt.Println("imgMap got:",imgMap)
		wg4Download.Add(1)
		go func(im map[string]string) {
			chSem <- 123
			DownloadImgWithClient(im["url"], im["filename"])
			<-chSem
			wg4Download.Done()
		}(imgMap)
	}

	wg4Download.Wait()
}

func GetImginfosFromPage(url string) {
	imginfos := GetPageImginfos(url)
	fmt.Println("imginfos=", imginfos)

	for _, infoMap := range imginfos {
		chImgMaps <- infoMap
		//fmt.Println("infoMap input:",infoMap)
	}
}
