package h_pkg

import (
	"strconv"
	"fmt"
)

/*
interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。
*/

type Human struct {
	Name string
	Age int
	Phone string
}

type Student struct{
	Human 
	School string
	Loan float32
}

type Employee struct {
	Human
	Company string
	Money float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}
//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}
//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Name,
		e.Company, e.Phone)
}

func Interfacetest()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	//定义Men类型的变量i
	var i Men
	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")
	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")
	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x{
		value.SayHi()
	}
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}


// Comma-ok断言
/*
Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
*/

type Element interface {

}

type List []Element
type Person struct{
	Name string
	Age int
}

func (p Person) String() string  {
	return "(name:" + p.Name + "--age:" + strconv.Itoa(p.Age) + ")"
}

func CommaOk()  {
	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{Name:"Dennis", Age:24}
	fmt.Printf("%#v\n", list[2])
	for index, element := range list {
		if val, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, val)
		} else if val, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, val)
		} else if val, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, val)
		}
	}

	// 这里有一点需要强调的是：element.(type)语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用comma-ok。
	for index, element := range list{
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}