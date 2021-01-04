package sync

import "sync"

func NewCounter() *Counter {
	return &Counter{}
}

type Counter struct {
	//添加一种互斥锁
	mu    sync.Mutex //私有的方法放在mu中 不要使用直接嵌入，这样会过度暴露不需要外界知道的方法
	value int
}

func (c *Counter) Inc() {
	//使用互斥锁
	//互斥对象在第一次使用后不能复制。
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
