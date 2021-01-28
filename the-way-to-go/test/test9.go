package main

import "fmt"

//Map
func main() {

	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.2
	mapCreated["key3"] = 7.8

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])

	//如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	//map 是 引用类型 的： 内存用 make 方法来分配。
	//map 的初始化：var map1 = make(map[keytype]valuetype)。
	//或者简写为：map1 := make(map[keytype]valuetype)。

	//不要使用new来创建map!
	//注意 如果你错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：

	//使用方法类型 作为值
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

	//用切片作为值
	//mp1 := make(map[int][]int)
	//mp2 := make(map[int]*[]int)

	//判断某个key是否存在
	//_, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false

	//删除某个key 如果 key1 不存在，该操作不会产生错误
	//delete(mp1, key1)

	//for range
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
	//注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的
	//map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
	//如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，
	//使用 sort 包再对切片排序，
	//然后可以使用切片的 for-range 方法打印出所有的 key 和 value。

	//map类型的切片 需要进行两次make
	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

}
