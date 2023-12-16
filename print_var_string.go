package main

import "fmt"

func main() {

	var name string = "Kodecloud"
	var user string = "Tamal"

	fmt.Print("welcome to ", name, ", ", user, "\n")

	var name1 string = "Hari"
	var grade int = 78
	fmt.Printf("Hey! %v you have scored %d \n", name1, grade)

	var firstname, lastname string = "tamal", "kundu"
	fmt.Println(firstname)
	fmt.Println(lastname)

	var (
		firstname1 string = "Tamal"
		age        int    = 32
	)
	fmt.Println(firstname1)
	fmt.Println(age)
}
