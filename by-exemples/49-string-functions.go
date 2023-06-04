package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	value := "test"
	p("Contains:", s.Contains(value, "es"))
	p("Count:", s.Count(value, "t"))
	p("HasPrefix:", s.HasPrefix(value, "te"))
	p("HasSuffix:", s.HasSuffix(value, "st"))
	p("Index:", s.Index(value, "e"))
	p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
}
