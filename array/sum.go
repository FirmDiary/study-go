package array

func Sum(numbers []int) (sum int) {
	sum = 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	// range 迭代数组 每次返回索引 与 值
	// 使用 _ 空白字符串 来忽略索引
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//func SumAll(numbersToSum ...[]int) (sums []int) {
//	lengthOfNumbers := len(numbersToSum)
//
//	//make 可以在创建切片的时候指定我们需要的长度和容量。
//	sums = make([]int, lengthOfNumbers)
//
//	for i, numbers := range numbersToSum {
//		sums[i] = Sum(numbers)
//	}
//	return
//}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		//append 函数，它能为切片追加一个新值。
		sums = append(sums, Sum(numbers))
	}
	return
}
