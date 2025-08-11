package main

import "fmt"

type Responce struct {
	Data    int
	Massage string
	Err     string
}

func main() {

	fmt.Println("total of tow number : ", addToNumber(40, 10)) // adding tow number function

	// total, msg := addProNumbers(40, 50, 60, 70, 80) // adding more numbers
	// fmt.Println(msg, total)

	responce := addProNumbers(40, 60, 100)
	if responce.Err == "Error" {
		return
	}
	fmt.Println(responce.Massage, responce.Data)

}

func addToNumber(num1 int, num2 int) int {
	total := num1 + num2
	return total
}

// func addProNumbers(nums ...int) (int, string) {

// 	total := 0
// 	for _, v := range nums {
// 		total += v

// 	}
// 	return total, "Total number is :"
// }

func addProNumbers(nums ...int) Responce {

	total := 0
	for _, v := range nums {
		total += v

	}
	return Responce{total, "Total numbers is :", "No Error"}
}
