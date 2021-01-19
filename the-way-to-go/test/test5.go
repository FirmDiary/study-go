package main

import "fmt"

//指针
func main() {
	var i1 = 4

	//使用&放到变量前面，就可以返回变量的内存地址
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	//如果想要调用指针intP，将*放到类型前面
	var intP *int
	intP = &i1                         //intP指向了i1的内存地址；它指向了i1的位置；它引用了变量i1；
	fmt.Printf("intP now is %p", intP) //%p 格式化指针
	fmt.Println()

	//一个指针变量可以指向任何一个值的内存地址。
	//在32位机器为4个字节，在64位机器为8个字节，与它指向的值的大小无关

	//当一个指针被声明但是未分配变量时，它是nil
	//一个指针变量的命名缩写一般为ptr

	//将*放在一个指针前面,叫做反引用,就可以得到这个指针上面的值
	fmt.Printf("intP now is %d", *intP) //%p 格式化指针

	s := "good bye"
	var p *string = &s
	*p = "再见" //更改指针内容.s也会随之被更改
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)

	//不能拿到一个常量或者一个数字文字的地址
	//const i = 5
	//prt := &i //cannot take the address of i
	//ptr2 := &10 //cannot take the address of 10
	//ptr3 := &"你好" //cannot take the address of "你好"
	str := "你好"
	prt4 := &str
	fmt.Printf("'str's pointer is %p\n", prt4)

	//移动指针指向字符串的字节数或数组的某个位置）是不被允许的
	//指针也可以指向另一个指针，并且可以进行任意深度的嵌套，导致你可以有多级的间接引用，但在大多数情况这会使你的代码结构不清晰

	//对空指针的反向引用是不合法的
	// invalid memory address or nil pointer dereference
	//var p2 *int = nil
	//*p2 = 0
}
