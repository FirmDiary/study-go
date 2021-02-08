package main

import "fmt"

//接口内定义了一系列方法和返回类型，但是不能有方法体和变量
//谁能够全部实现接口内的方法，谁就可以算是属于这个接口
//类型是隐式实现接口的
//类型可以实现多个接口

//定义一个接口
//type Namer interface {
//	Method1(params) return_type
//	Method2(params) return_type
//}
//（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，
//例如 Printer、Reader、Writer、Logger、Converter 等等。
//还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，
//此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。
//Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。
//不像大多数面向对象编程语言，在 Go 语言中接口可以有值，
//一个接口类型的变量或一个 接口值 ：var ai Namer，ai 是一个多字（multiword）数据结构，
//它的值是 nil。它本质上是一个指针，虽然不完全是一回事。指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误。

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

//Interface
func main() {

	//多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为。

	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

	//接口可以嵌套接口
	//type ReadWrite interface {
	//	Read(b Buffer) bool
	//	Write(b Buffer) bool
	//}
	//type Lock interface {
	//	Lock()
	//	Unlock()
	//}
	//type File interface {
	//	ReadWrite
	//	Lock
	//	Close()
	//}

	//类型判断
	//varI := varI.(T)
	//if v, ok := varI.(T); ok {  // checked type assertion
	//	Process(v)
	//	return
	//}
	//如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false，也没有运行时错误发生。
	//应该总是使用上面的方式来进行类型断言。

	//指针方法可以通过指针调用
	//值方法可以通过值调用
	//接收者是值的方法可以通过指针调用， 因为指针会首先被解引用
	//接收者是指针的方法不可以通过值调用， 因为存储在接口中的值没有地址
}
