package main

import "github.com/barnabasSol/DS/hashtable"

func main() {
	ht := hashtable.Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
	}
	for _, name := range list {
		ht.Insert(name)
	}
	println(ht.Search("BUTTERS"))
}
