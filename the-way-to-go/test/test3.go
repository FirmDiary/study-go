package main

import "fmt"

//给类型加别名
//就可以使用新类型操作int类型的数据
//但是新类型不会拥有原类型附带的方法
type TZ int

type ST string

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c is %d", c)

	var d ST = "你好"
	fmt.Printf("d is %s", d)
}
