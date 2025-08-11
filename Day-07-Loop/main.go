package main

import "fmt"

func main() {
	userList := []string{"vinul", "sewwandi", "sandaru", "avidu", "ravi"}

	for _, user := range userList {
		fmt.Println("User name is ", user)
	}

	for i := 0; i < len(userList); i++ {
		fmt.Println("User name is ", userList[i])

	}

}
