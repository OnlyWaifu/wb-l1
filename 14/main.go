package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeDefinition1()
	typeDefinition2()
}

func typeDefinition1() {
	fmt.Println("First method")
	arr := []any{"hi", 42, func() {}, struct{}{}, true, 45.6}
	for _, v := range arr {
		v := reflect.ValueOf(v)
		fmt.Printf("'%v' value has '%s' type\n", v, v.Kind().String())
	}
}

func typeDefinition2() {
	fmt.Println("Second method")
	var foo any = "bar"

	switch foo.(type) {
	case string:
		fmt.Println("foo is a string")
	case int:
		fmt.Println("foo is an int")
	case bool:
		fmt.Println("foo is a bool")
	case chan any:
		fmt.Println("foo is a chan")
	default:
		fmt.Println("foo is of some other type")
	}
}
