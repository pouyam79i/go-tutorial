package main

import (
	"fmt"
)

type Animal interface{
	Speek() string
}

type Dog struct{
	Grade int
}

type Cat struct{
	Grade int
}

func (d Dog) Work(){
	fmt.Println("Grade of dog:", d.Grade)
}

func (d Dog) Speek() string{
	return "Waaassshhhh"
}

func (c Cat) Speek() string{
	fmt.Println("cat in speek :", c.Grade)
	return "Meow"
}

func main()  {
	checkDefer()
	poodle := Animal(Dog{12})
	fmt.Println(poodle.Speek())
}

func checkDefer (){
	fmt.Println("This is a first line code of checkDefer")
	defer fmt.Println("This is a second line code of checkDefer")
	fmt.Println("This is a third line code of checkDefer")
	defer fmt.Println("Last in put of stack in defer")
}
