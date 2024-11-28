package main

import (
	"github.com/barnabasSol/DS/trie"
)

func main() {
	trie := trie.New()
	trie.Insert("barney")
	trie.Insert("ryan")
	trie.Insert("gosling")
	trie.Insert("drive")
	println(trie.Search(""))

}
