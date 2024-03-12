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
	assert.True(t, CompareMatrices(expectedResult, solvedMatrix[0]))

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

func TestGetMatrixColumn(t *testing.T) {
	matrix := [][]Cell{
		{{Value: 1}, {Value: 2}, {Value: 3}},
		{{Value: 4}, {Value: 5}, {Value: 6}},
	}

	column := getMatrixColumn(matrix, 1)
	expected := []Cell{{Value: 2}, {Value: 5}}
	if len(column) != len(expected) {
		t.Errorf("Expected column length %d, got %d", len(expected), len(column))
	}
	for i := range column {
		if column[i].Value != expected[i].Value {
			t.Errorf("Expected value %d at index %d, got %d", expected[i].Value, i, column[i].Value)
		}
	}
}

func TestCheckSide(t *testing.T) {
	solution := []Cell{{IsMarked: true}, {IsMarked: false}, {IsMarked: true}}
	if !checkSide(solution) {
		t.Error("Expected true, got false")
	}

	solution = []Cell{{IsMarked: true}, {IsMarked: true}, {IsMarked: true}}
	if checkSide(solution) {
		t.Error("Expected false, got true")
	}
}

func TestCheckIfTouchesBottomWall(t *testing.T) {
	solution := [][]Cell{
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: false}},
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: false}},
		{{IsMarked: true}, {IsMarked: false}, {IsMarked: true}},
	}

	if !checkIfTouchesBottomWall(solution) {
		t.Error("Expected true, got false")
	}

	solution = [][]Cell{
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: false}},
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: false}},
		{{IsMarked: true}, {IsMarked: true}, {IsMarked: true}},
	}

	if checkIfTouchesBottomWall(solution) {
		t.Error("Expected false, got true")
	}
}

func TestCheckIfTouchesRightWall(t *testing.T) {
	solution := [][]Cell{
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: true}},
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: true}},
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: false}},
	}

	if !checkIfTouchesRightWall(solution) {
		t.Error("Expected true, got false")
	}

	solution = [][]Cell{
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: true}},
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: true}},
		{{IsMarked: false}, {IsMarked: false}, {IsMarked: true}},
	}

	if checkIfTouchesRightWall(solution) {
		t.Error("Expected false, got true")
	}
}

func TestCombineMatrixes(t *testing.T) {
	matrix1 := [][]Cell{
		{{IsMarked: true}, {IsMarked: false}},
		{{IsMarked: false}, {IsMarked: true}},
	}
	matrix2 := [][]Cell{
		{{IsMarked: false}, {IsMarked: false}},
		{{IsMarked: false}, {IsMarked: true}},
	}
	expected := [][]Cell{
		{{IsMarked: false}, {IsMarked: false}},
		{{IsMarked: false}, {IsMarked: true}},
	}
	result := combineMatrixes(matrix1, matrix2)
	if !CompareMatrices(result, expected) {
		t.Error("Matrices are not combined correctly")
	}
}

func TestDuplicateMatrix(t *testing.T) {
	matrix := [][]Cell{
		{{IsMarked: true}, {IsMarked: false}},
		{{IsMarked: false}, {IsMarked: true}},
	}
	duplicate := DuplicateMatrix(matrix)
	if !CompareMatrices(matrix, duplicate) {
		t.Error("Duplicate matrix is not equal to the original matrix")
	}
}

func TestCompareMatrices(t *testing.T) {
	matrix1 := [][]Cell{
		{{Value: 1, IsMarked: true}, {Value: 2, IsMarked: false}},
		{{Value: 3, IsMarked: false}, {Value: 4, IsMarked: true}},
	}
	matrix2 := [][]Cell{
		{{Value: 1, IsMarked: true}, {Value: 2, IsMarked: false}},
		{{Value: 3, IsMarked: false}, {Value: 4, IsMarked: true}},
	}
	if !CompareMatrices(matrix1, matrix2) {
		t.Error("Matrices are not equal")
	}

	matrix3 := [][]Cell{
		{{Value: 1, IsMarked: true}, {Value: 2, IsMarked: false}},
		{{Value: 3, IsMarked: false}, {Value: 4, IsMarked: true}},
	}
	matrix4 := [][]Cell{
		{{Value: 1, IsMarked: true}, {Value: 2, IsMarked: false}},
		{{Value: 3, IsMarked: false}, {Value: 5, IsMarked: true}},
	}
	if CompareMatrices(matrix3, matrix4) {
		t.Error("Matrices are equal but shouldn't be")
	}
}

func TestPrintMatrix(t *testing.T) {
	matrix := [][]Cell{
		{{Value: 1, IsMarked: true}, {Value: 2, IsMarked: false}},
		{{Value: 3, IsMarked: false}, {Value: 4, IsMarked: true}},
	}
	PrintMatrix(matrix) // Just for visual inspection
}

func TestValidateInputMatrix(t *testing.T) {
	// Test case: Empty matrix
	emptyMatrix := [][]int{}
	if err := ValidateInputMatrix(&emptyMatrix); err == nil || err.Error() != "matrix is empty" {
		t.Errorf("Expected error: matrix is empty, got: %v", err)
	}

	// Test case: Single row matrix
	singleRowMatrix := [][]int{{1, 2, 3}}
	if err := ValidateInputMatrix(&singleRowMatrix); err == nil || err.Error() != "matrix is required not a single row or column" {
		t.Errorf("Expected error: matrix is required not a single row or column, got: %v", err)
	}

	// Test case: Single column matrix
	singleColumnMatrix := [][]int{{1}, {2}, {3}}
	if err := ValidateInputMatrix(&singleColumnMatrix); err == nil || err.Error() != "matrix is required not a single row or column" {
		t.Errorf("Expected error: matrix is required not a single row or column, got: %v", err)
	}

	// Test case: Valid matrix
	validMatrix := [][]int{{1, 2}, {3, 4}}
	if err := ValidateInputMatrix(&validMatrix); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestGetExampleResult(t *testing.T) {
	matrix := getTaskInit()
	solvedMatrix := SolveMatrix(matrix)
	assert.True(t, len(solvedMatrix) == 1)
}
