package main

import (
	"fmt"
	"os"

	"github.com/jgonis/anagramfinder/runemap"
	"github.com/jgonis/anagramfinder/wordlist"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		panic("Please provide a word to search for in the wordlist.")
	}
	inputWord := argsWithoutProg[0]
	inputRuneMap := runemap.NewRuneMap(inputWord)

	anagrams := wordlist.GetWordList(inputRuneMap)
	for _, word := range anagrams {
		fmt.Println(word)
	}
}
