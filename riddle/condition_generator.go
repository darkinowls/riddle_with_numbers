package riddle

const (
	minValue   = 1
	maxValue   = 3
	MatrixSize = 3
)

// GenerateAllMatrices generates all possible matrices of the specified size with numbers from 0 to 9.
func GenerateAllMatrices(size int) [][][]int {
	var matrices [][][]int
	var currentMatrix []int

	generateMatrixHelper(&matrices, &currentMatrix, 0, size)
	return matrices
}

// generateMatrixHelper recursively generates all possible matrices of the specified size with numbers from 0 to 9.
func generateMatrixHelper(matrices *[][][]int, currentMatrix *[]int, index, size int) {
	if index == size*size {
		m := reshapeToMatrix(*currentMatrix, size)
		tm := TranslateToCells(m)
		s := SolveMatrix(tm)
		if s == nil {
			return
		}
		*matrices = append(*matrices, m)
		return
	}

	for i := minValue; i <= maxValue; i++ {
		*currentMatrix = append(*currentMatrix, i)
		generateMatrixHelper(matrices, currentMatrix, index+1, size)
		*currentMatrix = (*currentMatrix)[:len(*currentMatrix)-1] // backtrack
	}
}

// reshapeToMatrix reshapes a 1D slice into a 2D slice of the specified size.
func reshapeToMatrix(matrix []int, size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			res[i][j] = matrix[i*size+j]
		}
	}
	return res
}
