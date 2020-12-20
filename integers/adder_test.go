package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

	//如果删除注释 「//Output: 6」，示例函数将不会执行。虽然函数会被编译，但是它不会执行。
}

// go test -v
// godoc -http=:6060
