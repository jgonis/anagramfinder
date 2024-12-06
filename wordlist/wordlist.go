package wordlist

import (
	"strings"

	"github.com/jgonis/anagramfinder/runemap"
)

func GetWordList(inputRuneMap runemap.RuneMap) []string {
	wordList := []string{}
	for _, word := range builtInWords {
		word = strings.ToLower(word)
		if inputRuneMap.CanCreateWord(word) {
			wordList = append(wordList, word)
		}
	}

	return wordList
}
