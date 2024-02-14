package main

import (
	"fmt"
	"math"
)

// Розробіть алгоритм вирішення задачі та реалізуйте його у вигляді
// програми мовою GOlang
// Зафарбуйте деяĸі ĸлітини таĸ, щоб у ĸожному рядĸу або стовпці
// не було чисел, що повторюються. Зафарбовані ĸлітини можуть стиĸатися одна з одною.
// Усі незафарбовані ĸлітини повинні
// з'єднуватися одна з одною сторонами по горизонталі або по
// вертиĸалі таĸ, щоб вийшов єдиний безперервний простір із
// незафарбованих ĸлітин.

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

// Алгритм
// 1. Проходимися по матриці по рядкам. Якщо значення в рядку повторюється, то йде роздвоєння через рекурсію:
// 		1) Нічого не змінюється
// 		2) У всії інштх значеннях забирається маркування
// 2. Проходимися по матриці по стовпцям. Те саме що й попередній крок
// 3.
//

func main() {
	initial := getTaskInit()
	result := SolveMatrix(initial)
	if CompareMatrices(result, getExampleResult()) {
		println("\nMatrices are IDENTICAL")
		return
	}
	println("\nMatrices are DIFFERENT")
}

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

func checkMatrixHasNoMarkedRowOrColumn(matrix [][]Cell) bool {
	// Check for marked rows
	for i := 0; i < len(matrix); i++ {
		rowMarked := true
		for j := 0; j < len(matrix[i]); j++ {
			if !matrix[i][j].IsMarked {
				rowMarked = false
				break
			}
		}
		if rowMarked {
			return false // Found a marked row
		}
	}

	// Check for marked columns
	for j := 0; j < len(matrix[0]); j++ {
		colMarked := true
		for i := 0; i < len(matrix); i++ {
			if !matrix[i][j].IsMarked {
				colMarked = false
				break
			}
		}
		if colMarked {
			return false // Found a marked column
		}
	}

	return true // No marked rows or columns found
}

func toggleMarkedWhereValue(matrix [][]Cell, value int8) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j].Value == value {
				matrix[i][j].IsMarked = !matrix[i][j].IsMarked
			}
		}
	}
}

func SolveMatrix(matrix [][]Cell, iStart int, jStart int, kStart int, lStart int) [][]Cell {
	uniqueValuesMap := make(map[int8]bool)

	// Check rows
	for i := iStart; i < len(matrix); i++ {
		countMap1 := make(map[int8]int)
		for j := jStart; j < len(matrix[i]); j++ {
			if matrix[i][j].IsMarked {
				continue
			}
			if countMap1[matrix[i][j].Value] > 0 {
				matrix[i][j].IsMarked = true
			} else {
				uniqueValuesMap[matrix[i][j].Value] = true
			}
			countMap1[matrix[i][j].Value]++
		}
	}

	// Check columns
	for k := kStart; k < len(matrix[0]); k++ {
		countMap2 := make(map[int8]int)
		for l := lStart; l < len(matrix); l++ {
			if matrix[l][k].IsMarked {
				continue
			}
			if countMap2[matrix[l][k].Value] > 0 {
				matrix[l][k].IsMarked = true
			} else {
				uniqueValuesMap[matrix[l][k].Value] = true
			}
			countMap2[matrix[l][k].Value]++
		}
	}

	if checkMatrixHasNoMarkedRowOrColumn(matrix) {
		return matrix
	}

	var uniqueValues []int8
	for key, _ := range uniqueValuesMap {
		uniqueValues = append(uniqueValues, key)
	}

	// Calculate the total number of combinations
	totalCombinations := int(math.Pow(2, float64(len(uniqueValues))))

	for i := 0; i < totalCombinations; i++ {
		// Toggle values to have unique combinations
		for j, value := range uniqueValues {
			if (i>>j)&1 == 1 { // Check if jth bit is set in i
				toggleMarkedWhereValue(matrix, value)
			}
		}

		printMatrix(matrix)

		if i == totalCombinations-1 || checkMatrixHasNoMarkedRowOrColumn(matrix) {
			return matrix
		}
	}

	return matrix
}

func printMatrix(matrix [][]Cell) {
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
