package main

//初始化一个玩家信息存储器
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//存储玩家信息
type InMemoryPlayerStore struct {
	store map[string]int
}

//记录赢
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//获取某个用户的分数
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
