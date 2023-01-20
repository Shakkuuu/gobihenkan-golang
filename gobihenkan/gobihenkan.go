package gobihenkan

import (
	"strings"
)

var Keyword = map[string]string{
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

func ChangeGobi(w string) string {
	var i bool
	for key, value := range Keyword {
		w = strings.ReplaceAll(w, key, value)
	}
	i = strings.HasSuffix(w, "のだ")
	if i {
		// fmt.Println(w)
		return w
	}

	o := strings.HasSuffix(w, "る")
	if o {
		w += "のだ"
		// fmt.Println(w)
		return w
	}

	s := strings.HasSuffix(w, "た")
	if s {
		w += "のだ"
		// fmt.Println(w)
		return w
	}

	w += "なのだ"
	// fmt.Println(w)
	return w
}
