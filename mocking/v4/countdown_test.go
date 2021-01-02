package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	//模拟sleep了四秒
	spySleeper := &SpySleeper{}

	//测试应该可以该通过，并且不再需要 4 秒
	Countdown(buffer, spySleeper)

	got := buffer.String()

	//反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 but got %d", spySleeper.Calls)

	}
}
