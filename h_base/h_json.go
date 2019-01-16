package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

type User2 struct {
	User string
	Name string `json:"nickname"`
	//Age      int
	//Birthday string
	//Sex      string
	//Email    string
	//Phone    string
}

/*结构体转json*/

func testStruct() {
	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	var user2 User2
	err = json.Unmarshal([]byte(data), &user2)
	if err != nil || user2 == (User2{}) {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(user2)
}

func testMap() {
	var mmp map[string]interface{}
	mmp = make(map[string]interface{})

	mmp["username"] = "user"
	mmp["age"] = 19
	mmp["sex"] = "man"

	data, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	fmt.Println("%s\n", string(data))

	var m map[string]interface{}
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println("ummarshal:", m)
}

func main() {
	testStruct()
	//testMap()
	fmt.Println("----")
}
