package main

import (
	"fmt"
)

//for
func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	//ASCII 编码的字符占用 1 个字节，既每个索引都指向不同的字符，
	//而非 ASCII 编码的字符（占有 2 到 4 个字节）
	//不能单纯地使用索引来判断是否为同一个字符

	// 1: 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值。
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}
	// 2:goto语句重写循环
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}

	gs2 := ""
	for i := 0; i < 25; i++ {
		gs2 += "G"
		fmt.Printf("%s\n", gs2)
	}

	//打印宽20高10的*
	w, h := 20, 10
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//for 条件语句 {}
	var i2 = 5

	for i2 >= 0 {
		i2--
		fmt.Printf("The variable i is now: %d\n", i)
	}

	//无限循环
	//如果for头部没有东西，那就认为是无限循环
	//for  {
	//}

	//无限循环的经典应用是服务器，用于不断等待和接受新的请求
	//for t, err = p.Token(); err == nil; t, err = p.Token() {
	//	...
	//}

	str3 := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str3 {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str4 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str4 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str4 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
	//每个 rune 字符和索引在 for-range 循环中是一一对应的。它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。

	//案例

	for i := 0; i < 5; i++ {
		var v int //每次都重新创建为0了
		fmt.Printf("%d ", v)
		v = 5
	}

	//无限循环
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	fmt.Println()
	//使用标签
	//for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，
	//即某一行第一个以冒号（:）结尾的单词
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				//continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置
				//通过continue 某个标签，达到跳出外层循环
				//如果直接使用continue，就会只跳出内层循环
				//如果直接使用break，就会终止所有循环了
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	//使用goto加上标签可以模拟for循环
	//但是这是不被鼓励的,用for将会更加方便
	i4 := 0
HERE:
	print(i4)
	i4++
	if i4 == 5 {
		return
	}
	goto HERE
}
