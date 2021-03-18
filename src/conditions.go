/*
Coded by Pouya Mohammadi
In this code I am trying to find out how conditions and loops work here
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	var a, b, c uint16
	a = 13
	b = 10
	c = 8

	// else if statement, remember the way is writing else if and else statement :)
	if a == 10 {
		fmt.Println("It is a")
	}else if b == 10 {
		a = 18
	}else {
		fmt.Println(c, " is C")
	}

	switch a {
	case 13:
		fmt.Println("This is 12")
	case 18:
		fmt.Println("This is 18")
	default:
		fmt.Println("It is nothing")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	var c1, c2, c3 chan int			// vars from a communication channel
	var i1, i2 int
	// select statement
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}

