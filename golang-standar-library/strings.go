package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fandi Hasnur", "Fandi"))
	fmt.Println(strings.Split("Fandi Hasnur", " "))
	fmt.Println(strings.ToLower("Fandi Hasnur"))
	fmt.Println(strings.ToUpper("Fandi Hasnur"))
	fmt.Println(strings.Trim("    Fandi Hasnur    ", " "))
	fmt.Println(strings.ReplaceAll("Fandi Hasnur Fandi Hasnur", "Fandi", "Alam"))
}
