package main

import (
	"fmt"
)

var loginToken string = "DGKDRANGDRBGSF24SD"

func main() {

	var address string
	username := "Vinul"

	address = "Wikum , Nadugala,Matarw"
	var isEmpty bool = true
	var avg float32 = 45.34563252424
	var avg2 float64 = 45.34563252424

	if isEmpty {
		fmt.Println("My name is ", username)
	}

	isEmpty = false

	if isEmpty {
		fmt.Println(address)
	}

	fmt.Println(avg)
	fmt.Println(avg2)
	fmt.Println(loginToken)

}
