package v1

import "math"

//go 语言中 interface resolution 是隐式的
type Shape interface {
	Area() float64
}

//可以使用保留字 struct 来定义自己的类型。
//一个通过 struct 定义出来的类型是一些已命名的域的集合，这些域用来保存数据

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

//float64 是形如 123.45 的浮点数
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Width) * 2
}
