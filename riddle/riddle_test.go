package riddle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateToCells(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]Cell{
		{{true, 1}, {true, 2}, {true, 3}},
		{{true, 4}, {true, 5}, {true, 6}},
		{{true, 7}, {true, 8}, {true, 9}},
	}

	output := TranslateToCells(input)

	assert.Equal(t, expected, output)
}

func TestSolveMatrix(t *testing.T) {
	// Test case with provided example matrix
	exampleMatrix := getExampleInit()
	expectedResult := GetExampleResult()
	solvedMatrix := SolveMatrix(exampleMatrix)
	assert.True(t, CompareMatrices(expectedResult, solvedMatrix))

	// Additional test cases can be added for different scenarios
}

func TestCheckIfUniqueWithinUnmarked(t *testing.T) {
	matrix := [][]Cell{
		{{false, 1}, {false, 2}, {false, 3}},
		{{false, 4}, {false, 5}, {false, 6}},
		{{false, 7}, {false, 8}, {false, 9}},
	}

	// Testing a matrix where all values are unique within unmarked cells
	assert.True(t, checkIfUniqueWithinUnmarked(matrix, 0, 0))

	// Testing a matrix where a value is repeated within unmarked cells
	matrix[0][1].Value = 1
	assert.False(t, checkIfUniqueWithinUnmarked(matrix, 0, 0))

	// Additional test cases can be added to cover various scenarios
}

func TestSolveMatrix_EmptyMatrix(t *testing.T) {
	// Testing when input matrix is empty
	var emptyMatrix [][]Cell
	solvedMatrix := SolveMatrix(emptyMatrix)
	assert.Nil(t, solvedMatrix)
}

func TestSolveMatrix_SingleColumn(t *testing.T) {
	// Testing with a single column matrix
	singleColumnMatrix := [][]Cell{
		{{false, 1}},
		{{false, 2}},
		{{false, 3}},
	}
	solvedMatrix := SolveMatrix(singleColumnMatrix)
	assert.Nil(t, solvedMatrix)
}

func TestSolveMatrix_AllSameValues(t *testing.T) {
	// Testing with a matrix where all values are the same
	allSameValuesMatrix := [][]Cell{
		{{false, 1}, {false, 1}, {false, 1}},
		{{false, 1}, {false, 1}, {false, 1}},
		{{false, 1}, {false, 1}, {false, 1}},
	}
	solvedMatrix := SolveMatrix(allSameValuesMatrix)
	assert.Nil(t, solvedMatrix)
}
