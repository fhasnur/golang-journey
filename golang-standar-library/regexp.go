package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`f([a-z]n)`)

	fmt.Println(regex.MatchString("fun"))
	fmt.Println(regex.MatchString("fan"))
	fmt.Println(regex.MatchString("fAn"))

	fmt.Println(regex.FindAllString("fun fan f4n fAn fen fzn f0n", 10))
}
