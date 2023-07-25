package main

import "fmt"

func countRectangles(_arrays [][]int) (numRectangles int) {
	// Make a copy of the arrays to avoid modifying the original
	arrays := make([][]int, len(_arrays))
	for i := range _arrays {
		arrays[i] = make([]int, len(_arrays[i]))
		copy(arrays[i], _arrays[i])
	}

	for startRow := 0; startRow < len(arrays); startRow++ {
		for startCol := 0; startCol < len(arrays[startRow]); startCol++ {
			if arrays[startRow][startCol] == 0 {
				continue
			}

			for endRow := len(arrays) - 1; endRow >= startRow; endRow-- {
				rectangleSize := 0
				for endCol := len(arrays[endRow]) - 1; endCol >= startCol; endCol-- {
					if arrays[endRow][endCol] == 0 {
						continue
					}

					if isRectangle(arrays, startRow, startCol, endRow, endCol) {
						numRectangles++
						cleanRectangle(arrays, startRow, startCol, endRow, endCol)

						// Skip the rest of the columns
						rectangleSize, endRow, endCol = endCol-startCol+1, startRow, startCol
					}
				}
				startCol += rectangleSize
			}
		}
	}
	return
}

func isRectangle(arrays [][]int, startRow, startCol, endRow, endCol int) bool {
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			if arrays[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func cleanRectangle(arrays [][]int, startRow, startCol, endRow, endCol int) {
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			arrays[i][j] = 0
		}
	}
}

func main() {
	arrays := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	numRectangles := countRectangles(arrays)
	fmt.Println(numRectangles)
}
