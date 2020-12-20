package iteration

const repeatCount = 5

func Repeat(character string, repeatCountGot int) string {

	if repeatCountGot == 0 {
		repeatCountGot = repeatCount
	}

	//:= 来声明和初始化变量
	//var 是显示声明
	var repeated string

	for i := 0; i < repeatCountGot; i++ {
		repeated += character
	}
	return repeated
}
