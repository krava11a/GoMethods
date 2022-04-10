package main

import "fmt"

func TestTypeAssertions() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s,ok)

	f, ok := i.(float64)
	fmt.Println(f,ok)

	i = 42
	z := i.(int)
	fmt.Println(z)

	s, ok = i.(string)
	fmt.Println(s,ok)


}
