/*
@Time : 2019-06-28 15:35
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"net/http"

	"github.com/micro/go-micro/web"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><body><h1>Hello World</h1></body></html>`)
}

func main() {

	//go test1()

	//go test2()

	go test3()

	select {}
}

func test1() {
	service := web.NewService(
		web.Name("helloworld"),
		web.Metadata(map[string]string{"cpu": "12%"}),
	)

	//cli := service.Client()
	//service.Init()
	//web.Metadata(map[string]string{"cpu": "12%"})

	service.HandleFunc("/", helloWorldHandler)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func test2() {
	service := web.NewService(
		web.Name("test.micro"),
	)

	service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, `<html><body><h1>micro</h1></body></html>`)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func test3() {
	service := micro.NewService(
		micro.Name("micro"),
		micro.Version("1.0.2"),
	)

	//fmt.Println(service.String())
	//service.Client().Call()
	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
