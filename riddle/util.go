package riddle

import (
	"errors"
	"fmt"
)

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

func IsInMatrixArray(matrix [][]Cell, matrixArray [][][]Cell) bool {
	for _, m := range matrixArray {
		if CompareMatrices(matrix, m) {
			return true
		}
	}
	return false
}

func ValidateInputMatrix(matrix *[][]int) error {
	if len(*matrix) == 0 {
		return errors.New("matrix is empty")
	}
	if len((*matrix)[0]) == 1 || len(*matrix) == 1 {
		return errors.New("matrix is required not a single row or column")
	}
	return nil
}
