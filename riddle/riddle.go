package riddle

// Розробіть алгоритм вирішення задачі та реалізуйте його у вигляді
// програми мовою GOlang
// Зафарбуйте деяĸі ĸлітини таĸ, щоб у ĸожному рядĸу або стовпці
// не було чисел, що повторюються. Зафарбовані ĸлітини можуть стиĸатися одна з одною.
// Усі незафарбовані ĸлітини повинні
// з'єднуватися одна з одною сторонами по горизонталі або по
// вертиĸалі таĸ, щоб вийшов єдиний безперервний простір із
// незафарбованих ĸлітин.

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Solve by pathfinding and combining 2 paths

func SolveMatrix(matrix [][]Cell) [][][]Cell {

	if len(matrix) == 0 || len(matrix[0]) == 1 || len(matrix) == 1 {
		return nil
	}

	// make way down from the top
	var solutionsDown [][][]Cell
	for i := 0; i < len(matrix[0]); i++ {
		solutionsDown = append(solutionsDown, makeWayDown(DuplicateMatrix(matrix), 0, i)...)
	}

	// make way from the left to the right
	var solutionsRight [][][]Cell
	for i := 0; i < len(matrix); i++ {
		solutionsRight = append(solutionsRight, makeWayRight(DuplicateMatrix(matrix), i, 0)...)
	}

	var solutions [][][]Cell = nil

	// combine the ways and check if it doesn't break the rules
	for _, solutionDown := range solutionsDown {
		for _, solutionRight := range solutionsRight {
			combined := combineMatrixes(solutionDown, solutionRight)
			if iterateMatrixAndCheckIfGood(combined) && !IsInMatrixArray(combined, solutions) {
				solutions = append(solutions, combined)
			}
		}
	}

	return solutions
}

func makeWayDown(originMatrix [][]Cell, initRow, initColumn int) (result [][][]Cell) {
	solutions := exploreMatrix(originMatrix, initRow, initColumn)
	for _, solution := range solutions {
		if checkIfTouchesBottomWall(solution) {
			result = append(result, solution)
		}
	}
	return
}

func makeWayRight(originMatrix [][]Cell, initRow, initColumn int) (result [][][]Cell) {
	solutions := exploreMatrix(originMatrix, initRow, initColumn)
	for _, solution := range solutions {
		if checkIfTouchesRightWall(solution) {
			result = append(result, solution)
		}
	}
	return
}

var directions = [4][2]int{
	{-1, 0}, // top
	{1, 0},  // bottom
	{0, -1}, // left
	{0, 1},  // right
}

func exploreMatrix(matrix [][]Cell, row, column int) (solutions [][][]Cell) {
	// Recursive function to explore cells
	var explore func([][]Cell, int, int)
	explore = func(matrix [][]Cell, row, column int) {
		matrix[row][column].IsMarked = false

		// Explore adjacent cells
		for dirIndex := 0; dirIndex < 4; dirIndex++ {
			dir := directions[dirIndex]
			// create new matrix to go the direction
			newRow, newColumn := row+dir[0], column+dir[1]

			// Check if the new position is within bounds else skip
			if newRow < 0 || newRow >= len(matrix) || newColumn < 0 || newColumn >= len(matrix[0]) {
				continue
			}

			// Check if the new position is not one of the previous positions
			if matrix[newRow][newColumn].IsMarked == false {
				continue
			}

			// Check if the new position has unique value to int the unmarked row
			isUnique := checkIfUniqueWithinUnmarked(matrix, newRow, newColumn)

			if isUnique {

				// create new dimension where it goes to the new position
				newMatrix := DuplicateMatrix(matrix)
				solutions = append(solutions, newMatrix)

				// go to the new position and discover path
				explore(newMatrix, newRow, newColumn)

				// go next direction
			}
		}
	}
	// start recursion
	explore(matrix, row, column)
	return
}

/////////////////////////////////////////////////////////////////////////////////////
// Check matrix by itself

func iterateMatrixAndCheckIfGood(matrix [][]Cell) bool {

	// find first unmarked
	var startCol int
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i].IsMarked == false {
			startCol = i
			break
		}
	}

	isGood := true
	visited := make(map[[2]int]bool) // Keep track of visited cells to avoid revisiting
	var iterate func(int, int)
	iterate = func(row int, col int) {
		if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) ||
			visited[[2]int{row, col}] || matrix[row][col].IsMarked == true {
			return
		}
		// Mark the cell as visited
		visited[[2]int{row, col}] = true
		// Perform check operation
		if checkIfUniqueWithinUnmarked(matrix, row, col) == false {
			isGood = false
			return
		}
		// Recursively call iterate for neighboring cells
		iterate(row+1, col) // Right
		iterate(row-1, col) // Left
		iterate(row, col+1) // Down
		iterate(row, col-1) // Up

	}
	iterate(0, startCol)
	return isGood
}

func checkIfUniqueWithinUnmarked(matrix [][]Cell, rowIndex int, columnIndex int) bool {
	value := matrix[rowIndex][columnIndex].Value
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == columnIndex || i == rowIndex {

				if j == columnIndex && i == rowIndex {
					continue
				}

				if matrix[i][j].Value == value && matrix[i][j].IsMarked == false {
					return false
				}

			}
		}
	}
	return true
}

func getMatrixColumn(matrix [][]Cell, columnNumber int) (firstColumn []Cell) {
	for i := 0; i < len(matrix); i++ {
		firstColumn = append(firstColumn, matrix[i][columnNumber])
	}
	return
}

func checkSide(solution []Cell) bool {
	for _, v := range solution {
		if v.IsMarked == false {
			return true
		}
	}
	return false
}

func checkIfTouchesBottomWall(solution [][]Cell) bool {
	hasBottom := false

	hasBottom = checkSide(solution[len(solution)-1])
	if hasBottom == false {
		return false
	}
	return true
}

func checkIfTouchesRightWall(solution [][]Cell) bool {
	hasRight := false
	hasRight = checkSide(getMatrixColumn(solution, len(solution[0])-1))
	if hasRight == false {
		return false
	}
	return true
}

func combineMatrixes(matrix1 [][]Cell, matrix2 [][]Cell) [][]Cell {
	combined := DuplicateMatrix(matrix1)
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[0]); j++ {
			if matrix2[i][j].IsMarked == false {
				combined[i][j].IsMarked = false
			}
		}
	}
	return combined
}
