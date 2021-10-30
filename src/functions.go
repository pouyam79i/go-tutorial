package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Println("It is some 2 :", addTowInt(10, 22))
	fmt.Println("It is add all:", addAll(12, 323, 23, 12, 32))
	n, l := FullName("Ali", "Moradi")
	fmt.Printf("FN: %v, Len: %v\n", n, l)
	n2, l2 := FullNameNakedReturn("Reza", "Moradi")
	fmt.Printf("FN: %v, Len: %v\n", n2, l2)
}

func doSomething() {
	fmt.Println("I am doing something")
}

func addTowInt(val1 int, val2 int) int {
	return val1 + val2
}

func addAll(vals ...int) int {
	sum := 0
	for i := range vals{
		sum += vals[i]
	}
	return sum
}

func FullName(firstName, lastName string) (string, int){
	fullName := firstName + " " + lastName
	length := len(fullName)
	return fullName, length
}

func FullNameNakedReturn(firstName, lastName string) (fullName string, length int){
	fullName = firstName + " " + lastName
	length = len(fullName)
	return
}
