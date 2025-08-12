package main

import (
	"fmt"
	"os"
)

func main() {
	fileWriter()
	fileReader()
}

func fileWriter() {

	file, err := os.Create("demo.txt")
	errorHandaler(err)
	contand := "Hi , How are you. My name is vinul. I'm form matara "
	len, err2 := file.WriteString(contand)
	errorHandaler(err2)
	fmt.Println("File contand size is :", len)
	defer file.Close()

}

func fileReader() {
	data, err := os.ReadFile("demo.txt")
	errorHandaler(err)
	fmt.Println(string(data))

}

func errorHandaler(err error) {
	if err != nil {
		panic(err)
	}
}
