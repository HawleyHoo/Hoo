package h_pkg

import (
	"fmt"
	"sort"
	"strconv"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
/*
我们没有实现HumanGroup的排序函数，所做的只是实现了三个函数（Len，Less和Swap），这个就是sort.Sort函数需要的全部信息。

我知道你很奇怪，你很想知道这个神奇之处是怎么实现的。实际上他的实现很简单，Sort包的排序函数接受任意类型的参数，只要他实现了Sort接口类型。

我们尝试了几种不同的利用接口类型作为参数的例子，这些例子利用接口类型达到了抽象数据类型的目的。

作者：Zuozuohao
链接：http://www.jianshu.com/p/dbd4e6b4900c
來源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test1()  {
	list := []int {1, 23, 65, 11, 0, 3, 233, 88, 99}
	fmt.Println("The list is: ", list)

	sort.Ints(list)
	fmt.Println("The sorted list is: ", list)


	group := HumanGroup{
		Human{Name:"Bart", Age:24},
		Human{Name:"Bob", Age:23},
		Human{Name:"Gertrude", Age:104},
		Human{Name:"Paul", Age:44},
		Human{Name:"Sam", Age:34},
		Human{Name:"Jack", Age:54},
		Human{Name:"Martha", Age:74},
		Human{Name:"Leo", Age:4},
	}

	//Let's print this group as it is
	fmt.Println("The unsorted group is:")
	for _, v := range group{
		fmt.Println(v)
	}

	//Now let's sort it using the sort.Sort function
	sort.Sort(group)

	//Print the sorted group
	fmt.Println("\nThe sorted group is:")
	for _, v := range group{
		fmt.Println(v)
	}

}

func (h Human) String() string {
	return "(name: " + h.Name + " - Age: "+strconv.Itoa(h.Age)+ " years)"
}

type HumanGroup []Human //HumanGroup is a type of slices that contain Humans

func (g HumanGroup) Len() int {
	return len(g)
}

func (g HumanGroup) Less(i, j int) bool {
	if g[i].Age < g[j].Age {
		return true
	}
	return false
}

func (g HumanGroup) Swap(i, j int){
	g[i], g[j] = g[j], g[i]
}


