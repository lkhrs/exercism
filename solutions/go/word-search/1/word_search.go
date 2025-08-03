package wordsearch

import (
	"fmt"
)

var directions = [][]int{
	// row, column
	{-1, -1}, // top left
	{-1, 0},  // top
	{-1, 1},  // top right
	{0, -1},  // center left
	{0, 1},   // center right
	{1, -1},  // bottom left
	{1, 0},   // bottom center
	{1, 1},   // bottom right
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	coords := make(map[string][2][2]int)
	for _, word := range words {
		for y, row := range puzzle {
			for x := range row {
			Search:
				for _, direction := range directions {
					xx, yy := x, y
					for i := 0; i < len(word); i, xx, yy = i+1, xx+direction[0], yy+direction[1] {
						if xx < 0 ||
							yy < 0 ||
							xx >= len(row) ||
							yy >= len(puzzle) ||
							puzzle[yy][xx] != word[i] {
							continue Search
						}
					}
					coords[word] = [2][2]int{{x, y}, {xx - direction[0], yy - direction[1]}}
				}
			}
		}
		if _, word := coords[word]; !word {
			err := fmt.Errorf("Can't find word \"%v\"", word)
			return nil, err
		}
	}
	return coords, nil
}
