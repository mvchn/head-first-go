package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(34))
	fmt.Println(reflect.TypeOf("head first go"))
}
