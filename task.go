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
// Solve by pathfinding

func SolveMatrix(matrix [][]Cell) [][]Cell {
	//startingPoints := matrix[0]
	//for i, _ := range startingPoints {
	//	solution := makeWay(DuplicateMatrix(matrix), 2)
	//	if solution != nil {
	//		return solution
	//	}
	//}
	solution := makeWay(DuplicateMatrix(matrix), 0, 4)
	if solution != nil {
		return solution
	}
	return nil
}

func makeWay(originMatrix [][]Cell, initRow, initColumn int) [][]Cell {
	// Initialize starting position
	directions := [][]int{
		{-1, 0}, // top
		{1, 0},  // bottom
		{0, -1}, // left
		{0, 1},  // right
	}
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

				if discMatrix[8][4].IsMarked == false {
					PrintMatrix(discMatrix)
				}

				// continue discovering path
				discMatrix = explore(discMatrix, row, column, dirIndex+1)

				// go next direction
			}
		}
		return discMatrix
	}

	explore(originMatrix, initRow, initColumn, 0)

	for _, solution := range solutions {
		if checkIfTouchesAllTheWalls(solution) {
			return solution
		}
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////
// Check matrix by itself

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

func checkIfTouchesAllTheWalls(solution [][]Cell) bool {
	hasLeft, hasRight, hasBottom := false, false, false

	hasBottom = checkSide(solution[len(solution)-1])
	if hasBottom == false {
		return false
	}
	hasLeft = checkSide(getMatrixColumn(solution, 0))
	if hasLeft == false {
		return false
	}
	hasRight = checkSide(getMatrixColumn(solution, len(solution[0])-1))
	if hasRight == false {
		return false
	}
	return true
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
