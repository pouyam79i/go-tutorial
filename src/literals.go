/*
Coded by Pouya Mohammadi
In this code I have learned about defining literals
plus showing hex, oct and bin is the same way as C and many more languages
 */

package main

import "fmt"

func main() {
	// my constants
	const myHex uint64 = 0XAB05C			// showing hexadecimal number
	const myOct uint64 = 0127301			// showing octal number
	const myBin uint64 = 0b10101001			// showing binary number
	const  myStr = "This is my str\nItis a literal\tWhich is a literal\n:)"
	// printing
	fmt.Println(myStr)
	fmt.Println("myHex: ", myHex, " myOct: ", myOct, " myBin: ", myBin)
}