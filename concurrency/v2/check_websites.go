package v1

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		//要告诉 Go 开始一个新的 goroutine，我们把一个函数调用变成 go 声明，
		//通过把关键字 go 放在它前面：go doSomething()
		go func(u string) {
			results[u] = wc(u)
		}(url) //传入当前url的副本，使其无法被改变
		//匿名函数
		//在声明的同时执行 —— 这就是匿名函数末尾的 () 实现的
		//维护对其所定义的词汇作用域的访问权 —— 在声明匿名函数时所有可用的变量也可在函数体内使用。

		//循环的每次迭代都会启动一个新的 goroutine，与当前进程（WebsiteChecker 函数）同时发生，每个循环都会将结果添加到 results map 中
	}

	time.Sleep(2 * time.Second)

	return results
}
