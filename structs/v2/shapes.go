package v1

//可以使用保留字 struct 来定义自己的类型。
//一个通过 struct 定义出来的类型是一些已命名的域的集合，这些域用来保存数据

type Rectangle struct {
	Width  float64
	Height float64
}

//float64 是形如 123.45 的浮点数
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Width) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
