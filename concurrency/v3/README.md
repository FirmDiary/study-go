- `goroutines`是Go的基本并发单元, 它让我们可以同时检查好多个网站
- `anonymous functions`(匿名函数), 我们用它来启动每个检查网站的并发进程
- `channels`, 用来组织和控制不同进程之间的交流，是我们能够避免`race condition`(竞争条件)的问题 
- `the race detector`（竞争探测器） 帮助我们调试并发代码的问题,
  我们无法准确控制每个 goroutine 写入结果 map 的时间，两个 goroutines 同一时间写入时程序将非常脆弱。
  Go 可以帮助我们通过其内置的 race detector 来发现竞争条件。要启用此功能，请使用 race 标志运行测试：go test -race


### 使程序加快
- 一种构建软件的敏捷方法，常常被错误地归属于 Kent Beck，即：
[让它运作，使它正确，使它快速]:http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast
  
- 「运作」是通过测试，「正确」是重构代码，而「快速」是优化代码以使其快速运行。一旦我们使程序可以正确运行，我们能做的就只有使它快速。很幸运，我们得到的代码已经被证明是可以运作的，并且不需要重构。在另外两个步骤执行之前，我们绝不应该试图「使它快速」，因为
[过早的优化是万恶之源]:http://wiki.c2.com/?PrematureOptimization —— Donald Knuth