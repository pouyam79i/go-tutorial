package main

import (
	"fmt"
)

type Student struct{
	FullName string
	Id int
	Grade int
}

func main()  {
	student1 := Student{"Pouya Mohammadi" ,9800000 ,5}
	fmt.Println(student1)
	fmt.Printf("%+v \n", student1)
}
