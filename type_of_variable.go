package main

import "fmt"

func main() {
	var grade int = 42
	var message string = "Hello world"
	var isCheck bool = true
	var amount float32 = 5466.64
	//Type of the variables.
	fmt.Printf("value of grade is = %v, and type is %T \n", grade, grade)
	fmt.Printf("value of message is = %v, and type is %T \n", message, message)
	fmt.Printf("value of isCheck is = %v, and type is %T \n", isCheck, isCheck)
	fmt.Printf("value of amount is = %v, and type is %T \n", amount, amount)
}
