package main 

import(
	"sort"
)

func IsAnagram(s , t string) bool {
	if len(s) != len(t) {
		return false 
	}

	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	}
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	}

	for i := range(sRunes) {
		if sRunes[i] != tRunes[i] {
			return false 
		}
	}
	return true
}
