package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i2                     int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

var inputReader *bufio.Reader
var input2 string
var err error

//读写
func main() {

	fmt.Println("Please enter you name")
	fmt.Scanln(&firstName, &lastName)
	//Scanln扫描输入文字，根据空格分隔依次存放到后续的参数内，换行结束

	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	inputReader = bufio.NewReader(os.Stdin)
	//函数返回一个新的带缓冲的 io.Reader 对象，它将从指定读取器（例如 os.Stdin）读取内容。
	fmt.Println("Please enter some input:")
	input, err = inputReader.ReadString('\n')
	//从输入中读取内容，直到碰到delim，并且和 delim连接起来一起输出
	fmt.Printf("Your name is %s", input)

}
