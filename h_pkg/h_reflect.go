package h_pkg

import (
	"reflect"
	"fmt"
)

type User struct {
	Id int
	Name string
	Addr string
}

func ReflectTest()  {
	var x float64 = 3.4
	v1 := reflect.ValueOf(x)
	fmt.Println("settability of v:", v1.CanSet()) // false

	v2 := reflect.ValueOf(&x)
	fmt.Println("settability of v:", v2.CanSet()) // false

	v := reflect.ValueOf(&x).Elem()
	fmt.Println("settability of v:", v.CanSet()) // true
	v.SetFloat(6.6)
	fmt.Println("x=", x)


	user := User{
		Id:12,
		Name:"jack",
		Addr:"hehehe",
	}


	//t := reflect.TypeOf(user)
	//v3 := reflect.ValueOf(user)
	/*
	reflect.ValueOf(&t)只是一个地址的值，通过.Elem()获取原始值对应的反射对象user
	所以上面两行等同于下面两行
	*/
	t := reflect.TypeOf(&user).Elem()
	v3 := reflect.ValueOf(&user).Elem()
	fmt.Printf("user :%#v\n", user)
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s : %v  type: %s\n", t.Field(k).Name, v3.Field(k).Interface(), v3.Field(k).Type())
	}

	name := v3.FieldByName("Name").String()
	fmt.Println("user.name:", name)
	v3.FieldByName("Name").SetString("rose")
	fmt.Println("set user.name:", user.Name)

	/*
	在面对类型时，需要区分 Type 和 Kind。
	前者表示真实类型（静态类型），后者表示其基础结构（底层类型）类别 -- 基类型。
	*/
	fmt.Println("kind:", t.Kind(), "type:",v3.Type())
	//kind: struct type: h_pkg.User

}

type user struct {
	name string
	age int
}
type manager struct {
	user
	title string
}

/*
只有在获取 结构体指针 的 基类型 后，才能遍历它的字段。
对于匿名字段，可用多级索引（按照定义顺序）直接访问。
FieldByName() 不支持多级名称，如有同名遮蔽，须通过匿名字段二次获取。
*/
func ReflectTest2() {
	var m manager
	t := reflect.TypeOf(&m)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // 如果是指针，则获取其所指向的元素
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(i, f.Name, f.Type, f.Offset)
		if f.Anonymous { // 输出匿名字段结构
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}

	t1 := reflect.TypeOf(m)
	name, _ := t1.FieldByName("name")
	fmt.Println(name.Name, name.Type)
	age := t.FieldByIndex([]int{0, 1}) // 按多级索引查找
	fmt.Println(age.Name, age.Type)
}

type A int
type B struct {
	A
}
func (A) av() {}
func (*A) ap() {}
func (B) bv() {}
func (*B) bp() {}

func ReflectTest3() {
	var b B
	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}
	for _, t2 := range s {
		fmt.Println(t2, ":")
		for i := 0; i < t2.NumMethod(); i++ {
			fmt.Println("m:", t2.Method(i).Name)
		}
	}
}

type User1 struct {
	Name string `field:"name" type:"varchar(50)"`
	Age  int `field:"age" type:"int"`
}
func ReflectTest4() {
	var u User1
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

/*辅助判断方法 Implements()、ConvertibleTo、AssignableTo() 都是运行期进行 动态调用 和 赋值 所必需的。*/

type X int

func (X) String() string  {
	return ""
}

func ReflectTest5()  {
	var a X
	t := reflect.TypeOf(a)
	// Implements 不能直接使用类型作为参数，导致这种用法非常别扭
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))
	it := reflect.TypeOf(0)
	fmt.Println(t.ConvertibleTo(it))
	fmt.Println(t.AssignableTo(st), t.AssignableTo(it))
}


