package main

import (
	"fmt"
	"time"
)

func main() {
	// time package usages
	precentTime := time.Now()
	fmt.Println("Today Time is :", precentTime)

	fmt.Println(precentTime.Format("01-02-2006 Monday")) //Get currunt date in sprcific format

	creatTime := time.Date(2026, time.December, 10, 12, 0, 0, 0, time.Local)

	fmt.Println(creatTime)
	fmt.Println("Create Time : ", creatTime.Format("2006-01-02 Monday "))

}
