package runemap

import "maps"

type RuneMap map[string]int

func NewRuneMap(word string) RuneMap {
	rm := make(RuneMap)
	for _, r := range word {
		rm[string(r)] += 1
	}
	return rm
}

func (rm RuneMap) CanCreateWord(word string) bool {
	wordRuneMap := NewRuneMap(word)
	return maps.Equal(rm, wordRuneMap)
}
