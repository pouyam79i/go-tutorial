package main

import (
	"fmt"
)

func main ()  {
	Arr := [5]int{1, 3, 5, 2, 4}
	for i := 0; i < len(Arr); i++{
		j := i - 1
		cur_val := Arr[i]
		for j >= 0 && Arr[j] > cur_val{
			Arr[j + 1] = Arr[j]
			j--
		}
		Arr[j + 1] = cur_val
	}		
	fmt.Println(Arr)
}