package main

import (
	"fmt"
	"os"
	"io"
)

func main()  {

	myContent := "This is my only string i want to save"
	
	file, err := os.Create("./text.txt") 
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, myContent)
	checkError(err)
	
	fmt.Printf("About writing: %v \n", ln)

}

func checkError(err error) {
	if err != nil{
		panic(err)
	}
	
}