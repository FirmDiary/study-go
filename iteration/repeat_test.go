package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("expected '%q' but got '%q'", want, got)
	}
}

func ExampleRepeat() {
	repeated := Repeat("Love", 2)
	fmt.Printf(repeated)
	// Output: LoveLove
}

//基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
//用 go test -bench=. 来运行基准测试。 (如果在 Windows Powershell 环境下使用 go test -bench=".")
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
