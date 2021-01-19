package main

import "fmt"

//switch
func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough //fallthrough关键字会让此分支不break,继续往下面执行
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	fmt.Println(monthToSeason(1))
	fmt.Println(monthToSeason(3))
	fmt.Println(monthToSeason(6))
	fmt.Println(monthToSeason(7))
	fmt.Println(monthToSeason(11))

}

func monthToSeason(month int) (season string) {
	switch month {
	case 3, 4, 5:
		season = "春天"
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		season = "夏天"
	case 9, 10, 11:
		season = "秋天"
	default:
		season = "冬天"
	}
	return
}
