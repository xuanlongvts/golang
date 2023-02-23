package main

import (
	"fmt"
	"os"
)

const (
	fileName = "localfile.data"
)

func main() {
	newData := []byte("I'm Anatoly\n")
	err := os.WriteFile(fileName, newData, os.ModePerm)
	if err != nil {
		fmt.Println("write file error: ", err)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read err: ", err)
	}
	fmt.Println(string(data))

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Open err: ", err)
	}
	defer f.Close()

	if _, err := f.WriteString("New data that wasn't there originally \n"); err != nil {
		fmt.Println("write file WriteString error: ", err)
	}

	data, err = os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read err: ", err)
	}
	fmt.Println(string(data))
}
