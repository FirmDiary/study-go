package main

import "fmt"

var i = 5
var str = "ABC"

type Person2 struct {
	name string
	age  int
}

//空接口
type Any interface{}

//Interface
func main() {

	var val Any
	val = i
	fmt.Printf("val has the value: %v\n", val)

	val = str
	fmt.Printf("val has the value: %v\n", val)

	val = Person2{name: "Azal", age: 23}
	fmt.Printf("val has the value: %v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person2:
		fmt.Printf("Type pointer to Person2 %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
