package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = []string{"Red", "Blue", "Yello"}

	colors = append(colors, "Green")
	fmt.Println(colors)

	// colors := append(colors[1:len(colors)])
	// fmt.Println(molors)

	colors = append(colors[2:])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 5
	numbers[1] = 2
	numbers[2] = 10
	numbers[3] = 5
	numbers[4] = 12
	numbers = append(numbers, 6)
	sort.Ints(numbers)
	fmt.Println(numbers)
	fmt.Println(len(numbers))

}
