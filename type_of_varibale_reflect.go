package main

import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 42
	var message string = "Hello World Tamal Kundu"
	// Type of the datatype.
	fmt.Printf("grades is %v, type of grades is %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("message is %v, type of message is %v \n", message, reflect.TypeOf(message))
}
