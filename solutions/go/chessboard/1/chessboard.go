package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedSquares := 0
	for _, v := range cb[file] {
		if v {
			occupiedSquares++
		}
	}
	return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
// Rank = index. Run through each File and return a count of indexes which are true
func CountInRank(cb Chessboard, rank int) int {
	occupiedSquares := 0
	if rank <= 8 && rank >= 1 {
		for i := range cb {
			for i, v := range cb[i] {
				if i == rank-1 && v {
					occupiedSquares++
				}
			}
		}
	}
	return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0
	for i := range cb {
		for range cb[i] {
			totalSquares++
		}
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalSquares := 0
	for i := range cb {
		for _, i := range cb[i] {
			if i {
				totalSquares++
			}
		}
	}
	return totalSquares
}
