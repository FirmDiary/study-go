package main

import "fmt"

//切片
func main() {

	//定义一个数组
	//必须指定长度
	//v := [3]int{1,2,3}

	//切片是没长度的数组
	//定义一个切片
	//v1 := []int{1,2,3}
	//0<=len(v1)<=cap(v1)  cap()可以获取切片的最大容量
	//v2 := []int{}  空切片为nil 长度len(v2)=0
	//切片是一个引用类型
	//切片可以从数组初始化
	//var v3 []int = v[:]  [start:end] 都可以省略代表等于整个数组
	//v3 = &v[:]

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // 2是引入的，而5没有

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	// grow the slice
	slice1 = slice1[:cap(slice1)]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//创建一个匿名函数
	//函数可以传递给参数,这是函数式编程的一大特定
	sum := func(a []int) (s int) { //通过定义返回数据,使代码更加整洁易读
		for i := 0; i < len(a); i++ {
			s += a[i]
		}
		return
	}
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(sum(arr[:]))

	sum1 := []int{}
	fmt.Printf("%d\n", sum1)
	//make的语法 func make([]T, len, cap)，其中 cap 是可选参数。
	//创建一个初始切片，len设置为10，此时cap(sum2)=len(sum2)=10
	sum2 := make([]int, 10)
	fmt.Printf("%d\n", sum2)
	sum3 := make([]int, 50, 100)
	fmt.Printf("%d\n", sum3)
	sum4 := new([100]int)[:50]
	fmt.Printf("%d\n", sum4)
	//new (T) 为每个新的类型 T 分配一片内存，
	//初始化为 0 并且返回类型为 * T 的内存地址：
	//这种方法 返回一个指向类型为 T，值为 0 的地址的指针， 它适用于值类型如数组和结构体；
	//它相当于 &T{}。
	//make(T) 返回一个类型为 T 的初始值，
	//它只适用于 3 种内建的引用类型：切片、map 和 channel
	//new 函数分配内存，make 函数初始化

	//var v []int = make([]int, 10, 50)
	//v := make([]int, 10, 50)
	//分配了一个有50个int值的数组，并且创建了一个长度为10，容量为50的切片v,v指向数组的前十个元素

	s := make([]byte, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = s[2:4]
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s1 := []byte{'a', 'z', 'a', 'l'}
	s2 := s1[2:]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%d", s3)

	fmt.Println()
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Printf("%d", items)
	for i := range items {
		items[i] *= 2
	}
	fmt.Printf("%d", items)
}
