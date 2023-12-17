package main

import "fmt"

func main() {

	var nam string
	var age int
	fmt.Print("enter name and age : ")
	count, err := fmt.Scanf("%s %d", &nam, &age)
	fmt.Println("count: ", count)
	fmt.Println("error: ", err)
	fmt.Print("name: ", nam, " age: ", age)
}
