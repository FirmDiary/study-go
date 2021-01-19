package main

import "fmt"

//func main() {
//	var a int
//	var b int32
//	a = 15
//	b = a + a // 编译错误  cannot use a + a (type int) as type int32 in assignment
//	b = b + 5 // 常量是可以隐士转换的, 因为 5 是常量，所以可以通过编译
//}

func main() {
	var n int16 = 34
	var m int32

	m = int32(n) //必须通过显示转换

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
