package main

import (
	"fmt"

	"github.com/Shakkuuu/gobihenkan-golang/gobihenkan"
)

func main() {
	var word string
	fmt.Scanln(&word)
	w := gobihenkan.ChangeGobi(word)
	fmt.Println(w)
}
