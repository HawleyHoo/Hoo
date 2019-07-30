package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

var visits = expvar.NewInt("visits")
var stats = expvar.NewMap("tcp")

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	stats.Add("requests", 2)
	stats.Add("requests_failed", 3)

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[0:])
}

func main01() {
	fmt.Println(stats)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1818", nil)
	/*
		http://localhost:1818/debug/vars
		http://localhost:1818/golang
	*/
}

func kvFunc(kv expvar.KeyValue) {
	fmt.Println(kv.Key, kv.Value)
}

func main() {
	inerInt := int64(10)
	pubInt := expvar.NewInt("Int")
	pubInt.Set(inerInt)
	pubInt.Add(2)

	inerFloat := 1.2
	pubFloat := expvar.NewFloat("Float")
	pubFloat.Set(inerFloat)
	pubFloat.Add(0.1)

	inerString := "hello"
	pubString := expvar.NewString(inerString)
	pubString.Set(inerString)

	pubMap := expvar.NewMap("Map").Init()
	pubMap.Set("Int", pubInt)
	pubMap.Set("Float", pubFloat)
	pubMap.Set("String", pubString)
	pubMap.Do(kvFunc)
	pubMap.Add("Int", 1)
	pubMap.Add("NewInt", 123)
	pubMap.AddFloat("Float", 0.5)
	pubMap.AddFloat("NewFloat", 0.9)
	pubMap.Do(kvFunc)

	expvar.Do(kvFunc)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	err := http.ListenAndServe(":1818", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
