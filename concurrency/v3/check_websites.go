package v1

type WebsiteChecker func(string) bool

type result struct {
	string //当不需要命名值的时候，可不命名，是匿名的
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		//要告诉 Go 开始一个新的 goroutine，我们把一个函数调用变成 go 声明，
		//通过把关键字 go 放在它前面：go doSomething()
		go func(u string) {
			// send statement
			//当我们迭代 urls 时，不是直接写入 map，而是使用 send statement 将每个调用 wc 的 result 结构体发送到 resultChannel。
			//这使用 <- 操作符，channel 放在左边，值放在右边
			resultChannel <- result{u, wc(u)}
		}(url) //传入当前url的副本，使其无法被改变
		//匿名函数
		//在声明的同时执行 —— 这就是匿名函数末尾的 () 实现的
		//维护对其所定义的词汇作用域的访问权 —— 在声明匿名函数时所有可用的变量也可在函数体内使用。

		//循环的每次迭代都会启动一个新的 goroutine，与当前进程（WebsiteChecker 函数）同时发生，每个循环都会将结果添加到 results map 中
	}

	//下一个 for 循环为每个 url 迭代一次。
	//我们在内部使用 receive expression，它将从通道接收到的值分配给变量
	for i := 0; i < len(urls); i++ {
		// receive expression
		//这也使用 <- 操作符，但现在两个操作数颠倒过来：
		//现在 channel 在右边，我们指定的变量在左边
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

//通过将结果发送到通道，我们可以控制每次写入 results map 的时间，确保每次写入一个结果。
//虽然 wc 的每个调用都发送给结果通道，但是它们在其自己的进程内并行发生，因为我们将结果通道中的值与接收表达式一起逐个处理一个结果。
//我们已经将想要加快速度的那部分代码并行化，同时确保不能并发的部分仍然是线性处理。
//我们使用 channel 在多个进程间通信
