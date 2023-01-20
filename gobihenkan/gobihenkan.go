package gobihenkan

import (
	"fmt"
	"strings"
)

var keyword = map[string]string{
	"まだ": "まだなのだ",
	"した": "したのだ",
	"った": "ったのだ",
	"です": "なのだ",
	"よう": "ようなのだ",
	"な":  "ななのだ",
}

func changegobi(w string) {
	for key, value := range keyword {
		w = strings.ReplaceAll(w, key, value)
	}
	fmt.Println(w)
}
