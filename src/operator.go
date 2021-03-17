/*
Coded by Pouya Mohammadi
In this program I am going to test the Operators that exist in Go
 */
package main

import "fmt"

func main() {
	var a, b int8;
	a = 12
	b = 9

	//	Arithmetic Operators like + - * / % ++ --
	var c = a + b
	fmt.Print("c");

	// Relation Operator like == >= <= != > <
	if a > b {
		fmt.Print("It is ture flag 1")
	}

	// Logical Operator like && || !
	if a > 20 && b < a {
		fmt.Println("It is true flag 2")
	}

	// Bitwise Operator like & | ^ << >>
	c = a | b
	fmt.Printf("This is or Operator %d", c)

	// Miscellaneous Operators are * and &, we use this one with pointers
	fmt.Println(&a)

}