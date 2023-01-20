package gobihenkan

import (
	"fmt"
	"strings"
)

var keyword = map[string]string{
	"まだ":   "まだなのだ",
	"した":   "したのだ",
	"った":   "ったのだ",
	"です":   "なのだ",
	"よう":   "ようなのだ",
	"します":  "するのだ",
	"される":  "されるのだ",
	"されます": "されるのだ",
	"られる":  "れるのだ",
	"ます":   "ますのだ",
}

func ChangeGobi(w string) {
	var i bool
	for key, value := range keyword {
		w = strings.ReplaceAll(w, key, value)
	}
	i = strings.HasSuffix(w, "のだ")
	if i {
		fmt.Println(w)
		return
	}

	o := strings.HasSuffix(w, "る")
	if o {
		w += "のだ"
		fmt.Println(w)
		return
	}

	s := strings.HasSuffix(w, "た")
	if s {
		w += "のだ"
		fmt.Println(w)
		return
	}

	w += "なのだ"
	fmt.Println(w)
}
