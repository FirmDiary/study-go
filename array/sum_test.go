package array

import "testing"

func TestSum(t *testing.T) {

	//t.Run("Collection of 5 numbers", func(t *testing.T) {
	//    numbers := [5]int{1, 2, 3, 4, 5}
	//
	//    got := Sum(numbers)
	//    want := 15
	//
	//    if got != want {
	//        t.Errorf("got %d want %d given, %v", got, want, numbers)
	//        //%v 默认输入格式 适合展示数组
	//    }
	//})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

//go test --cover 查看测试覆盖率
