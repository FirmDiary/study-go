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
