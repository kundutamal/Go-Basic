package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String conversion.
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T \n", err, err)
}
