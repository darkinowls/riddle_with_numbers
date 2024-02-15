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
		IsMarked: false,
		Value:    value,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	initial := getExampleInit()
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
// Solve by filtering

func SolveMatrix(matrix [][]Cell) [][]Cell {
	solutions := calculateAllPossibleMatrixSolutions(matrix)
	filteredSolutions := filterMatrixesThatTouchAllTheWalls(solutions)
	for _, s := range filteredSolutions {
		PrintMatrix(s)
		if iterateMatrixAndCountCells(s) == calculateUnmarkedCells(s) {
			return s
		}
	}
	return nil
}

func calculateAllPossibleMatrixSolutions(matrix [][]Cell) (solutions [][][]Cell) {

	initMatrix := getInitMatrixWithMarks(matrix)

	var iterateMatrix func(matrix [][]Cell, rowStartIndex, columnStartIndex int)
	iterateMatrix = func(matrix [][]Cell, rowStartIndex, columnStartIndex int) {
		for i := rowStartIndex; i < len(matrix); i++ {
			for j := columnStartIndex; j < len(matrix[0]); j++ {

				newMatrix := DuplicateMatrix(matrix)

				if newMatrix[2][0].IsMarked == false {
					// i never saw this print
					PrintMatrix(newMatrix)
				}

				newMatrix = toggleCellsInCrossWithTheSameValue(newMatrix, i, j)

				iterateMatrix(newMatrix, i, j+1)

			}
		}
		solutions = append(solutions, matrix)

	}
	iterateMatrix(initMatrix, 0, 0)
	return
}

/////////////////////////////////////////////////////////////////////////////////////
// Check matrix by itself

func getInitMatrixWithMarks(matrix [][]Cell) [][]Cell {
	solution := DuplicateMatrix(matrix)
	for i := 0; i < len(solution); i++ {
		for j := 0; j < len(solution[0]); j++ {
			if solution[i][j].IsMarked == true {
				continue
			}
			switchOtherCellsInCrossWithTheSameValue(solution, i, j) // continue
		}
	}
	return solution
}

func iterateMatrixAndCountCells(matrix [][]Cell) (count int) {

	// find first unmarked
	var startCol int
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i].IsMarked == false {
			startCol = i
			break
		}
	}

	visited := make(map[[2]int]bool) // Keep track of visited cells to avoid revisiting
	var iterate func(int, int)
	iterate = func(row int, col int) {
		if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || visited[[2]int{row, col}] {
			// Mark the cell as visited
			visited[[2]int{row, col}] = true
			// Perform operations on the cell here (e.g., count it)
			count++
			// Recursively call iterate for neighboring cells
			iterate(row+1, col) // Right
			iterate(row-1, col) // Left
			iterate(row, col+1) // Down
			iterate(row, col-1) // Up
		}
	}
	iterate(0, startCol)
	return
}

func calculateUnmarkedCells(matrix [][]Cell) (count int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j].IsMarked == false {
				count++
			}
		}
	}
	return
}

func switchOtherCellsInCrossWithTheSameValue(matrix [][]Cell, rowIndex int, columnIndex int) {
	value := matrix[rowIndex][columnIndex].Value
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == columnIndex || i == rowIndex {
				if j == columnIndex && i == rowIndex {
					continue
				}
				if matrix[i][j].Value == value {
					matrix[i][j].IsMarked = true
				}
			}
		}
	}
}

func toggleCellsInCrossWithTheSameValue(matrix [][]Cell, rowIndex int, columnIndex int) [][]Cell {
	value := matrix[rowIndex][columnIndex].Value
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == rowIndex || j == columnIndex {

				//if matrix[0][2].IsMarked == false && matrix[2][0].IsMarked == false {
				//	PrintMatrix(matrix)
				//}

				if i == rowIndex && j == columnIndex {
					matrix[i][j].IsMarked = false // itself in the cross
				} else if matrix[i][j].Value == value {
					matrix[i][j].IsMarked = true // others in the cross
				}
			}
		}
	}
	return matrix
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

func filterMatrixesThatTouchAllTheWalls(matrixes [][][]Cell) (filteredMatrixes [][][]Cell) {
	for _, m := range matrixes {
		if checkIfTouchesAllTheWalls(m) {
			filteredMatrixes = append(filteredMatrixes, m)
		}
	}
	return
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
