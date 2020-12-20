package array

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//在 Go 中不能对切片使用等号运算符。你可以写一个函数迭代每个元素来检查它们的值。
	//但是一种比较简单的办法是使用 reflect.DeepEqual，它在判断两个变量是否相等时十分有用。
	//if got != want {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want '%v', got '%v'", want, got)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
