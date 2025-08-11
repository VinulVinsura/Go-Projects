package main

import "fmt"

func main() {

	coures := make(map[string]string)
	nameList := make(map[int]string)

	coures["JS"] = "JAVA-SCRIPT"
	coures["TS"] = "TYPR-SCRIPT"
	coures["JV"] = "JAVA"

	nameList[1] = "vinul"
	nameList[2] = "ravi"
	nameList[4] = "sewwandi"

	fmt.Println("Couser Map is : ", coures)
	fmt.Println("User Map is : ", nameList)

	fmt.Println(nameList[1], " is Love ", nameList[4])

	delete(coures, "JS")
	fmt.Println("Couser Map is : ", coures)

}
