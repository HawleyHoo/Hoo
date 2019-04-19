package main

import (
	//	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Mall struct {
	name string
	cat  []Catagory
}

type Catagory struct {
	id   string
	name string
	//spec   string
	price string
	desc  string
	//subCat []*SubCatagory
}

//type SubCatagory struct {
//	id        int64
//	name      string
//	link      string
//	detailCat []*DetailCatagory
//}
//
//type DetailCatagory struct {
//	id    int64
//	name  string
//	link  string
//	goods map[string]interface{}
//}

var Jd = Mall{name: "山姆"}

var topCatagoryStart = regexp.MustCompile(`[[:space:]]*\<div[[:space:]]+class="result_innerBox clear"[[:space:]]+id="result_innerBox"\>`)
var topCatagoryEnd = regexp.MustCompile(`[[:space:]]*\<div[[:space:]]+class="mod_searchMoreBtn"[[:space:]]+id="mod_searchMoreBtn"\>`) // <div class="mod_searchMoreBtn" id="mod_searchMoreBtn">

var topCatagoryFetch = regexp.MustCompile(`[[:space:]]*.*\<span\>(?P<topC>.*)\</span\>`)
var subCatagoryFetch = regexp.MustCompile(``)

var detailCatagoryStart = regexp.MustCompile(`[[:space:]]*\<dt\>\<a[[:space:]]+href=\"//(?P<cat2link>.*)\"[[:space:]]+target="_blank"\>(?P<cat2name>[\p{Han}]+)\</a\>\</dt\>`)

//`<a href="//e.jd.com/products/5272-10941.html" target="_blank">网络原创</a>`
var detailCatagoryFetch = regexp.MustCompile(`[[:space:]]*.*\<a[[:space:]]+href=\"//(?P<cat2link>.*)\"[[:space:]]+target="_blank"\>(?P<cat2name>[\p{Han}]+)\</a\>`)

func dumpJdCatagory(mall *Mall) {
	fmt.Println(mall.name)
	for _, c := range Jd.cat {
		//fmt.Println(c)
		fmt.Printf("list: %+v\n", c)
		//for _, sc := range c.subCat {
		//	fmt.Printf("SubCatagory: %s, Link: %s\n", sc.name, sc.link)
		//	for _, dc := range sc.detailCat {
		//		fmt.Printf("DetailCatagory: %s. Link: %s\n", dc.name, dc.link)
		//	}
		//}
	}
}

func main() {
	//resp, err := http.Get("http://www.jd.com/allSort.aspx")
	resp, err := http.Get("http://list.samsclub.cn/search/c142101-1/?tp=2240.166041.0.0.0.M!EhZxr-10-FL5so&tps=x0.10375y0.05224&ti=M!EhZxr-10-FL5so_FG51")
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	//	fmt.Println(string(body))

	file, err := os.Create("sam_list.html")
	if err != nil {
		panic(err.Error())
	}

	_, err = file.Write(body)
	if err != nil {
		panic(err.Error())
	}

	Jd.cat = make([]Catagory, 0, 10000)
	var top = false
	//var sub = false
	//var detail = false
	cat := Catagory{}
	//var subCat *SubCatagory
	//var detailCat *DetailCatagory
	fmt.Println(top)
	s := strings.Split(string(body), "\n")
	for _, line := range s {
		if samSkuBegin.MatchString(line) {
			fmt.Println("match result_innerBox")
			top = true
			cat.id = samSkuBegin.FindStringSubmatch(line)[1]

			//	//sub = false
			//	detail = false
		}
		/*if top == true {
			if topCatagoryFetch.MatchString(line) {
				sub = true

					//fmt.Println(topCatagoryFetch.FindStringSubmatch(line)[1])
					//	cat = &Catagory{name: topCatagoryFetch.FindStringSubmatch(line)[1]}
					//	cat.subCat = make([]*SubCatagory, 40, 100)

				cat = new(Catagory)
				cat.name = topCatagoryFetch.FindStringSubmatch(line)[1]
				cat.subCat = make([]*SubCatagory, 0, 100)
				Jd.cat = append(Jd.cat, cat)
				//fmt.Println("Catagory")
				//fmt.Println(cat)
			}
		}

		if sub == true {
			if detailCatagoryStart.MatchString(line) {

					//fmt.Println(detailCatagoryStart.FindStringSubmatch(line)[1])
					//fmt.Println(detailCatagoryStart.FindStringSubmatch(line)[2])
					//	subCat = &SubCatagory{name: detailCatagoryStart.FindStringSubmatch(line)[2], link: detailCatagoryStart.FindStringSubmatch(line)[1]}
					//	subCat.detailCat = make([]*DetailCatagory, 50, 100)

				subCat = new(SubCatagory)
				subCat.name = detailCatagoryStart.FindStringSubmatch(line)[2]
				subCat.link = detailCatagoryStart.FindStringSubmatch(line)[1]
				subCat.detailCat = make([]*DetailCatagory, 0, 100)
				cat.subCat = append(cat.subCat, subCat)
				//fmt.Println("SubCatagory")
				//fmt.Println(subCat)
				detail = true
			}
		}*/
		if top {
			//fmt.Println("samFetch", samFetch.MatchString(line), line)

			//fmt.Println("Catagory")
			if samFetch.MatchString(line) {
				cat.name = samFetch.FindStringSubmatch(line)[2]
				fmt.Printf("fetch name:%+v \n", samFetch.FindStringSubmatch(line))
				//cat.subCat = make([]*SubCatagory, 0, 100)
			} else {

				//fmt.Println(line)
			}

			if samFetchPrice.MatchString(line) {
				fmt.Printf("fetch price:%+v \n", samFetchPrice.FindStringSubmatch(line))
				cat.price = samFetchPrice.FindStringSubmatch(line)[1] + samFetchPrice.FindStringSubmatch(line)[2]
			}

			if samFetchDesc.MatchString(line) {
				//fmt.Printf("fetch desc:%+v \n", samFetchDesc.FindStringSubmatch(line))
				//cat.desc = samFetchDesc.FindStringSubmatch(line)[1]
			}
			fmt.Println("line", line)
		}

		if samSkuEnd.MatchString(line) {
			fmt.Println("match result_innerBox end")
			Jd.cat = append(Jd.cat, cat)
			top = false
		}
	}

	dumpJdCatagory(&Jd)
	//fmt.Println(s)
	//fmt.Printf("%d bytes has been write to jd_list.html", n)

}
