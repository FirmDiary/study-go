package main

import (
	"fmt"
	"io"
	"net/http"
)

//用 writer 把问候发送到我们测试中的缓冲区
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//func main() {
//    Greet(os.Stdout, "Yanying")
//}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

//运行程序并访问 http://localhost:5000。你会看到你的 greeting 函数被使用了
func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
