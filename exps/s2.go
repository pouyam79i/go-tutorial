package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str1, _ := reader.ReadString('\n')
	str2, _ := reader.ReadString('\n')
	var word1 string = strings.TrimSpace(str1)
	var word2 string = strings.TrimSpace(str2)
	// fmt.Printf("Str1 was: %c - Str2 was: %c", word1[0], word2[2])
	if word1[0] == word1[2] || word1[0] == word2[2] {
		fmt.Print("YES")
	} else if word2[0] == word1[2] || word2[0] == word2[2] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
