package v1

import "fmt"

//type Stringer interface {
//	String() string
//}

//重写Bitcoin的string方法
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Go 允许从现有的类型创建新的类型。
//语法是 type MyName OriginalType
type Bitcoin int //比特币

//在go中，如果一个符合（例如变量，类型，函数等）是以小写符号开头
//那么它在定义它的包之外就是私有的
type Wallet struct {
	balance Bitcoin //私有的
}

//在 Go 中，当调用一个函数或方法时，参数会被复制。
//使用 w *Wallet 就把指针指向Wallet, 而不是创建一个副本
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
