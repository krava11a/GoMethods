package main

import "fmt"

type myInterface interface {
	Method()
}

type tStruct struct {
	S string
}

func (t *tStruct) Method() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)

}

func TestInterfaceValuseWithNil() {
	var i myInterface
	var t *tStruct
	describeMyInterface(i)
	//i.Method()
	i = t
	describeMyInterface(i)
	i.Method()

	i = &tStruct{"Vasia"}
	describeMyInterface(i)
	i.Method()

}

func describeMyInterface(m myInterface) {
	fmt.Printf("(%v, %T)\n", m, m)
}
