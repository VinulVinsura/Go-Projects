package main

import (
	"fmt"
)

func main() {

	// var userList = []string{}

	// userList = append(userList, "vinul")
	// userList = append(userList, "sewwandi")
	// userList = append(userList, "ravidu")

	// fmt.Println("User List is :", userList)

	// slices.Reverse(userList)
	// fmt.Println("User List is :", userList)

	// Remove element in slice using index
	cousers := []string{"java ,", "js ,", "go ,", "c# ,", "spring ,", "angular"}
	fmt.Println("This my corse list : ", cousers)
	var index int

	fmt.Print("Enter the remove couser index : ") //4
	fmt.Scan(&index)

	cousers = append(cousers[:index-1], cousers[index:]...)
	fmt.Println("Updated Course List is : ", cousers)

}
