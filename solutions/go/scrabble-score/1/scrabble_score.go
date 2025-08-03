package scrabble

import "strings"

func Score(word string) (score int) {
	values := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	for _, letter := range strings.ToUpper(word) {
		for points, scoringLetters := range values {
			if strings.ContainsRune(scoringLetters, letter) {
				score += points
			}
		}
	}
	return
}
