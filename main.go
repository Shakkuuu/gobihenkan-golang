package main

import (
	"fmt"

	"github.com/Shakkuuu/gobihenkan-golang/gobihenkan"
)

func main() {
	var word string
	fmt.Scanln(&word)
	gobihenkan.ChangeGobi(word)
}
