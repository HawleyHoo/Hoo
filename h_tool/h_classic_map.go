package main

import "fmt"

// 人员档案
type Profile struct {
	Name    string // 名字
	Age     int    // 年龄
	Married bool   // 已婚
}

func simpleHash(str string) (ret int) {

	// 遍历字符串中的每一个ASCII字符
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[i]

		// 将字符的ASCII码相加
		ret += int(c)
	}

	return
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}
	buildIndex(list)
	queryData("张三", 30)
}

// 查询键
type classicQueryKey struct {
	Name string // 要查询的名字
	Age  int    // 要查询的年龄
}

// 计算查询键的哈希值
func (c *classicQueryKey) hash() int {
	// 将名字的Hash和年龄哈希合并
	return simpleHash(c.Name) + c.Age*1000000
}

// 创建哈希值到数据的索引关系
var mapper = make(map[int][]*Profile)

// 构建数据索引
func buildIndex(list []*Profile) {

	// 遍历所有的数据
	for _, profile := range list {

		// 构建数据的查询索引
		key := classicQueryKey{profile.Name, profile.Age}

		// 计算数据的哈希值, 取出已经存在的记录
		existValue := mapper[key.hash()]

		// 将当前数据添加到已经存在的记录切片中
		existValue = append(existValue, profile)

		// 将切片重新设置到映射中
		mapper[key.hash()] = existValue
	}
}

// 根据条件查询数据
func queryData(name string, age int) {

	// 根据给定查询条件构建查询键
	keyToQuery := classicQueryKey{name, age}

	// 计算查询键的哈希值并查询, 获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]

	// 遍历结果集合
	for _, result := range resultList {

		// 与查询结果比对, 确认找到打印结果
		if result.Name == name && result.Age == age {
			fmt.Println("result:", result)
			return
		}
	}

	// 没有查询到时, 打印结果
	fmt.Println("no found")

}
