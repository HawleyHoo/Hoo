package main

import (
	"eshop/model"
	"eshop/util"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type SamSku struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	SkuID    string        `bson:"SkuID"`
	Catagory string        `bson:"Catagory"`
	Name     string        `bson:"Name"`
	Price    string        `bson:"Price"`
	Desc     string        `bson:"Desc"`
}

//
var samSkuBegin = regexp.MustCompile(`[[:space:]]*\<div[[:space:]]+class="mod_searchPro jsModSearfhPro"[[:space:]]+prodid=\"(?P<skuid>.*)\"`)
var samSkuEnd = regexp.MustCompile(`[[:space:]]*\<div[[:space:]]+class="proBorder"\>\</div\>`)

//`<a href="//item.samsclub.cn/77606975" target="_blank">MM黑鳕鱼块750gx1袋.</a>`
//`<a href="//cms.samsclub.cn/sale/izSq" target="_blank">隐私政策</a>`
//var samFetch = regexp.MustCompile(`[[:space:]]*.*\<a[[:space:]]+href=\"//(?P<cat2link>.*)\"[[:space:]]+target="_blank"\>(?P<cat2name>[\p{Han}]+)\</a\>`)
var samFetch = regexp.MustCompile(`[[:space:]]*.*\<a[[:space:]]+href=\"//(?P<cat2link>.*)\"[[:space:]]+target="_blank"\>(?P<skuname>.*)\</a\>`)

// <p class="proPrice clear"><em><b class="b1">¥ </b>96</em></p>
var samFetchPrice = regexp.MustCompile(`[[:space:]]*.*\<p[[:space:]]+class="proPrice clear"\>\<em\>\<b[[:space:]]+class="b1"\>(?P<priceUnit>.*)\</b\>(?P<price>.*)\</em\>\</p\>`)

var samFetchDesc = regexp.MustCompile(`(?P<cat2name>[\p{Han}]+)`)

func main() {
	//resp, err := http.Get("http://www.jd.com/allSort.aspx")
	urlsr := "http://list.samsclub.cn/search/c142009-1/?tp=2240.142002.0.0.0.M!F!Ttc-10-FL5so&tps=x0.28202y0.02709&ti=M!F!Ttc-10-FL5so_NUCQ"
	resp, err := http.Get(urlsr)
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

	var top = false
	samsku := SamSku{}

	s := strings.Split(string(body), "\n")
	for _, line := range s {
		if samSkuBegin.MatchString(line) {
			fmt.Println("match result_innerBox")
			top = true
			samsku.SkuID = samSkuBegin.FindStringSubmatch(line)[1]
			samsku.Catagory = "干果"

		}

		if top {
			//fmt.Println("samFetch", samFetch.MatchString(line), line)

			if samFetch.MatchString(line) {
				samsku.Name = samFetch.FindStringSubmatch(line)[2]
				//fmt.Printf("fetch name:%+v \n", samFetch.FindStringSubmatch(line))
			}

			if samFetchPrice.MatchString(line) {
				//fmt.Printf("fetch price:%+v \n", samFetchPrice.FindStringSubmatch(line))
				samsku.Price = samFetchPrice.FindStringSubmatch(line)[1] + samFetchPrice.FindStringSubmatch(line)[2]
			}

			//if samFetchDesc.MatchString(line) {
			//fmt.Printf("fetch desc:%+v \n", samFetchDesc.FindStringSubmatch(line))
			//cat.desc = samFetchDesc.FindStringSubmatch(line)[1]
			//}
			//fmt.Println("line", line)
		}

		if samSkuEnd.MatchString(line) {
			fmt.Println("match result_innerBox end")
			insertSamSku(&samsku)

			top = false
		}
	}
}

const samskuCollectionName = "SamSku"

func insertSamSku(o *SamSku) error {
	model.GetDbSessionInstance()
	session, col := model.GetDbSessionInstance().GetCollection(samskuCollectionName)
	defer session.Close()
	o.ID = bson.NewObjectId()
	err := col.Insert(o)
	if err != nil {
		util.LogInfo("sam sku insert err %s", err.Error())
	}
	return err
}
