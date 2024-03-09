package main

import (
	"fmt"
	"golang-basic/database"
	_ "golang-basic/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
