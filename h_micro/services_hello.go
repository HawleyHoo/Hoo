/*
@Time : 2019-06-28 17:15
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	service.Init()

}
