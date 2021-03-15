/*
Coded by Pouya Mohammadi
In this code I am going to use all kind of data types
which can be used in Go
 */

package src

import "fmt"

// Main fucntion of my app
func main() {
	var myBoolean bool
	var myUint8bit uint8
	//var my_uint_64bit uint64
	//var my_int_16bit int16
	//var my_int_32bit int32
	//var my_float_32bit float32
	//var my_float_64bit float64
	var my_complexNum_64bit complex64
	//var my_complexNum_128bit complex128
	//var my_uint_8bit_byte byte
	//var my_uint_32bit_rune rune
	//var my_uint_32or64bit uint
	//var my_int_32or64bit int
	//var my_pointer_address uintptr

	undeclearedType := 14.5; // Now it is float
	fmt.Println(undeclearedType)

	myUint8bit = 2
	fmt.Println(myUint8bit)

	myBoolean = true
	if myBoolean {
		fmt.Println("It is true")
	}

	my_complexNum_64bit = 13 + 13i
	fmt.Println(my_complexNum_64bit)




}