package main

import (
	"fmt"
	"sort"
)

func main()  {
	map1 := make(map[string]int)
	map1["A"] = 3
	map1["C"] = 1
	map1["M"] = 2
	map1["D"] = 4
	fmt.Println(map1)

	keys := make([]string, len(map1))
	i := 0
	for key := range map1{
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	for i := range keys{
		fmt.Printf("Key: %s - Value: %d\n", keys[i], map1[keys[i]])
	}

}
