package main

import "fmt"

type User struct {
	id       float32
	fullName string
	username string
	password string
}

// func setUserDetails(user *User) {
// 	user.id = 10
// 	user.fullName = "vinul vinsura"
// 	user.username = "vinu"
// 	user.password = "435er"
// }

func main() {

	// fmt.Println("Main func start")
	// var person User
	// fmt.Println(person)
	// setUserDetails(&person)
	// fmt.Println(person)

	var ptr *int
	number := 20
	fmt.Println(ptr)
	ptr = &number
	fmt.Println(ptr)

	*ptr = 50
	fmt.Println(ptr)
	fmt.Println(number)

}
