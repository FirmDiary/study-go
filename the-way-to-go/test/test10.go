package main

import (
	"fmt"
	"strings"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

//Struct
func main() {

	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	//type myStruct struct { i int }
	//var v myStruct    // v是结构体类型变量
	//var p *myStruct   // p是指向一个结构体类型变量的指针
	//v.i
	//p.i

	//初始化一个结构体实例（一个结构体字面量：struct-literal）的更简短和惯用的方式如下：
	//&struct1{a, b, c} 是一种简写，底层仍然会调用 new ()，这里值的顺序必须按照字段顺序来写
	mss := &struct1{10, 15.5, "Chris"}
	fmt.Println(mss)
	//或者
	var msss struct1
	msss = struct1{10, 15.5, "Chris"}
	fmt.Println(msss)

	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)

	//类型可以有tag 通过reflect 能够得到tag
	//type TagType struct { // tags
	//	field1 bool   "An important answer"
	//	field2 string "The name of the thing"
	//	field3 int    "How much there are"
	//}
	//ttType := reflect.TypeOf(tt)
	//ixField := ttType.Field(ix)
	//fmt.Printf("%v\n", ixField.Tag)

	//类型里可以匿名，并且可以嵌套类型
	//type innerS struct {
	//	in1 int
	//	in2 int
	//}
	//
	//type outerS struct {
	//	b    int
	//	c    float32
	//	int  // anonymous field  此时结构就是名字，每个类型中单个类型只能存在一个匿名
	//	innerS //anonymous field
	//}

	//可以为类型定义方法
	//可以把类型看作是面向对象里的类
	//
	//格式 类型就是接收者，
	//func(recv receiver_type) methodName(parameter_list) (return_value_list){....}
	//为接收者提供一个方法，这个方法是对接收者使用的 可以使用recv.methodName调用方法
	//接收者可以理解为面向对象里的this与self,指向类自己
	//如果不需要使用接收者的方法或值，可以用_忽略

	//类型和方法必须定义在同一个包里面

	//如果想要改变接收者里面的值，那么就要在接收者类型前面加上*，在他的指针类型上面定义方法

	//类型A中的值类型是另一个类型B时，这个类型可以直接使用B类型的方法，也可以在类型A中覆写类型B的方法。
	//内嵌将一个已存在类型的字段和方法注入到了另一个类型里：
	//匿名字段上的方法 “晋升” 成为了外层类型的方法。当然类型可以有只作用于本身实例而不作用于内嵌 “父” 类型上的方法，
	//可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。

	//聚合或者组合，一个具名字段类型
	//内嵌， 一个匿名字段类型

	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())

	//可以将两个类型以匿名的方式嵌入到一个新的类型中去，实现’多重继承‘

	//在 Go 中，类型就是类（数据和关联的方法）。Go 拥有类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。
	//在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 组件编程（Component Programming）。
}

//可把类型看出面向对象语言中的对象
//将类型设置为私有 就只可以通过工厂方法去建立类型
//type matrix struct {
//	//...
//}
//

//工厂方法建立类型
//func NewMatrix(params) *matrix {
//	m := new(matrix) // 初始化 m
//	return m
//}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
