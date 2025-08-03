package isogram

import "strings"

func IsIsogram(word string) bool {
	isIso := true
	for _, letter := range word {
		if letter == '-' || letter == ' ' {
			continue
		}
		lowerWord := strings.ToLower(word)
		lowerLetter := strings.ToLower(string(letter))
		count := strings.Count(lowerWord, lowerLetter)
		if count > 1 {
			isIso = false
		}
	}
	return isIso
}
