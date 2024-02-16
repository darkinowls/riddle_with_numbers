package main

import (
	"fmt"
)

// Розробіть алгоритм вирішення задачі та реалізуйте його у вигляді
// програми мовою GOlang
// Зафарбуйте деяĸі ĸлітини таĸ, щоб у ĸожному рядĸу або стовпці
// не було чисел, що повторюються. Зафарбовані ĸлітини можуть стиĸатися одна з одною.
// Усі незафарбовані ĸлітини повинні
// з'єднуватися одна з одною сторонами по горизонталі або по
// вертиĸалі таĸ, щоб вийшов єдиний безперервний простір із
// незафарбованих ĸлітин.

// /////////////////////////////////////////////////////////////////////////////////////////////////
// DATA STRUCTURE

type Cell struct {
	IsMarked bool
	Value    int8
}

func NewCell(value int8) Cell {
	return Cell{
		IsMarked: true,
		Value:    value,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	initial := getTaskInit()
	result := SolveMatrix(initial)
	PrintMatrix(result)

	//if CompareMatrices(result, getExampleResult()) {
	//	println("\nMatrices are IDENTICAL")
	//	return
	//}
	//println("\nMatrices are DIFFERENT")
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// INPUT / OUTPUT

func getTaskInit() [][]Cell {
	matrix := [][]Cell{
		{NewCell(1), NewCell(1), NewCell(4), NewCell(3), NewCell(4), NewCell(1), NewCell(3), NewCell(2), NewCell(2)},
		{NewCell(1), NewCell(1), NewCell(2), NewCell(3), NewCell(2), NewCell(1), NewCell(3), NewCell(2), NewCell(2)},
		{NewCell(3), NewCell(2), NewCell(1), NewCell(4), NewCell(3), NewCell(3), NewCell(2), NewCell(1), NewCell(3)},
		{NewCell(4), NewCell(3), NewCell(4), NewCell(2), NewCell(3), NewCell(1), NewCell(1), NewCell(2), NewCell(4)},
		{NewCell(4), NewCell(2), NewCell(1), NewCell(1), NewCell(2), NewCell(3), NewCell(3), NewCell(4), NewCell(1)},
		{NewCell(2), NewCell(2), NewCell(3), NewCell(3), NewCell(4), NewCell(4), NewCell(4), NewCell(1), NewCell(2)},
		{NewCell(2), NewCell(3), NewCell(3), NewCell(1), NewCell(3), NewCell(2), NewCell(2), NewCell(4), NewCell(1)},
		{NewCell(4), NewCell(4), NewCell(2), NewCell(1), NewCell(3), NewCell(1), NewCell(2), NewCell(3), NewCell(3)},
		{NewCell(4), NewCell(4), NewCell(2), NewCell(1), NewCell(1), NewCell(1), NewCell(2), NewCell(3), NewCell(3)},
	}
	return matrix
}

func getExampleInit() [][]Cell {
	matrix := [][]Cell{
		{NewCell(4), NewCell(2), NewCell(4), NewCell(8)},
		{NewCell(8), NewCell(6), NewCell(6), NewCell(8)},
		{NewCell(4), NewCell(2), NewCell(6), NewCell(6)},
		{NewCell(2), NewCell(2), NewCell(6), NewCell(6)},
	}
	return matrix
}

func getExampleResult() [][]Cell {
	matrix := [][]Cell{
		{Cell{true, 4}, Cell{false, 2}, Cell{false, 4}, Cell{false, 8}},
		{Cell{false, 8}, Cell{false, 6}, Cell{true, 6}, Cell{true, 8}},
		{Cell{false, 4}, Cell{true, 2}, Cell{true, 6}, Cell{true, 6}},
		{Cell{false, 2}, Cell{true, 2}, Cell{true, 6}, Cell{true, 6}},
	}
	return matrix
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Solve by pathfinding and intersection

func SolveMatrix(matrix [][]Cell) [][]Cell {
	startingPointsDown := len(matrix[0])
	startingPointsRight := len(matrix)

	var solutionsDown [][][]Cell
	var solutionsRight [][][]Cell

	// make way down from the top
	for i := 0; i < startingPointsDown; i++ {
		solutionsDown = append(solutionsDown, makeWayDown(DuplicateMatrix(matrix), 0, i)...)
	}

	// make way from the left to the right
	for j := 0; j < startingPointsRight; j++ {
		solutionsRight = append(solutionsRight, makeWayRight(DuplicateMatrix(matrix), j, 0)...)
	}

	// combine both ways
	// 1. find the intersection and filter by it
	var filteredSolutions [][][]Cell
	for _, solutionDown := range solutionsDown {
		for _, solutionRight := range solutionsRight {
			if checkIfHasIntersections(solutionDown, solutionRight) {

				filteredSolutions = append(filteredSolutions, solutionDown, solutionRight)
			}
		}
	}

	// 2. combine the ways and check if it doesn't break the rules
	for _, solutionDown := range solutionsDown {
		for _, solutionRight := range solutionsRight {
			combined := combineMatrixes(solutionDown, solutionRight)
			if iterateMatrixAndCheckIfGood(combined) {
				return combined
			}
		}
	}

	return nil
}

var directions = [4][2]int{
	{-1, 0}, // top
	{1, 0},  // bottom
	{0, -1}, // left
	{0, 1},  // right
}

func makeWayDown(originMatrix [][]Cell, initRow, initColumn int) [][][]Cell {
	// Initialize starting position

	var solutions [][][]Cell

	// Recursive function to explore cells
	var explore func([][]Cell, int, int, int) [][]Cell
	explore = func(matrix [][]Cell, row, column, startingDirIndex int) [][]Cell {
		matrix[row][column].IsMarked = false

		discMatrix := matrix

		// Explore adjacent cells
		for dirIndex := startingDirIndex; dirIndex < 4; dirIndex++ {
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
				discMatrix = explore(newMatrix, newRow, newColumn, 0)

				// go next direction
			}
		}
		return discMatrix
	}

	explore(originMatrix, initRow, initColumn, 0)

	var result [][][]Cell
	for _, solution := range solutions {
		if checkIfTouchesBottom(solution) {
			result = append(result, solution)
		}
	}
	return result
}

func makeWayRight(originMatrix [][]Cell, initRow, initColumn int) [][][]Cell {
	// Initialize starting position

	var solutions [][][]Cell

	// Recursive function to explore cells
	var explore func([][]Cell, int, int, int) [][]Cell
	explore = func(matrix [][]Cell, row, column, startingDirIndex int) [][]Cell {
		matrix[row][column].IsMarked = false

		discMatrix := matrix

		// Explore adjacent cells
		for dirIndex := startingDirIndex; dirIndex < 4; dirIndex++ {
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
				discMatrix = explore(newMatrix, newRow, newColumn, 0)

				// go next direction
			}
		}
		return discMatrix
	}

	explore(originMatrix, initRow, initColumn, 0)

	var result [][][]Cell
	for _, solution := range solutions {
		if checkIfTouchesRight(solution) {
			result = append(result, solution)
		}
	}
	return result
}

/////////////////////////////////////////////////////////////////////////////////////
// Check matrix by itself

func checkIfHasIntersections(m1 [][]Cell, m2 [][]Cell) bool {
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0]); j++ {
			if m1[i][j].IsMarked == false && m2[i][j].IsMarked == false {
				return true
			}
		}
	}
	return false

}

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
		if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || visited[[2]int{row, col}] || matrix[row][col].IsMarked == true {
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

func checkIfTouchesBottom(solution [][]Cell) bool {
	hasBottom := false

	hasBottom = checkSide(solution[len(solution)-1])
	if hasBottom == false {
		return false
	}
	return true
}

func checkIfTouchesRight(solution [][]Cell) bool {
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

/////////////////////////////////////////////////////////////////////////////////////////////////////
// Utils

func DuplicateMatrix(matrix [][]Cell) [][]Cell {
	n := len(matrix)
	m := len(matrix[0])
	duplicate := make([][]Cell, n)
	data := make([]Cell, n*m)
	for i := range matrix {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func PrintMatrix(matrix [][]Cell) {
	println()
	for _, row := range matrix {
		for _, cell := range row {
			if cell.IsMarked {
				fmt.Printf("*%-4d ", cell.Value)
			} else {
				fmt.Printf("%-5d ", cell.Value)
			}
		}
		fmt.Println()
	}
}

func CompareMatrices(matrix1, matrix2 [][]Cell) bool {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		return false // Matrices have different dimensions
	}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			if matrix1[i][j].Value != matrix2[i][j].Value || matrix1[i][j].IsMarked != matrix2[i][j].IsMarked {
				return false // Cells at position (i, j) are different
			}
		}
	}

	return true // Matrices are identical
}
