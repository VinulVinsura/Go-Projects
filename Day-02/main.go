package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Convertion int values
	name := "123"

	numConv, err := strconv.ParseInt(name, 0, 64)

	if err != nil {
		fmt.Println("Error is throw")
		fmt.Println(err)

	} else {
		fmt.Println(numConv)
	}

	// Convertion int values

}
