package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func httpGet() {
	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

/*
 http.Post()的 contentType 默认格式 application/x-www-form-urlencoded
*/
func httpPostForm() {
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func MakeParams(params url.Values, appKey string) (params_str, sign_str string) {
	var s, p string
	var keys []string
	b := bytes.Buffer{}
	b.WriteString(appKey)
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, v := range keys {
		b.WriteString(v)
		b.WriteString(params.Get(v))
	}
	p = b.String()
	b.WriteString(appKey)
	s = b.String()
	p = strings.TrimRight(p, "&")
	return p, s
}

func MakeParams2(params map[string][]string, appKey string) (params_str, sign_str string) {
	var s, p string
	var keys []string
	b := bytes.Buffer{}
	b.WriteString(appKey)
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, v := range keys {
		b.WriteString(v)
		b.WriteString(params[v][0])
	}
	p = b.String()
	b.WriteString(appKey)
	s = b.String()
	p = strings.TrimRight(p, "&")
	return p, s
}

/*作者：scloudrun
来源：CSDN
原文：https://blog.csdn.net/mingzhehaolove/article/details/51861510
版权声明：本文为博主原创文章，转载请附上博文链接！*/
