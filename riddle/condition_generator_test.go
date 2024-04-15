package riddle

import (
	"reflect"
	"testing"
)

func TestGenerateAllMatrices(t *testing.T) {
	// Test case: size 2 matrix
	hasMatrix := [][]int{
		{2, 1},
		{1, 2},
	}

	matrices := GenerateAllMatrices(2)
	if len(matrices) < 20 {
		t.Errorf("Expected at least 20 matrices, got %d", len(matrices))
	}
	for _, m := range matrices {
		if reflect.DeepEqual(m, hasMatrix) {
			return
		}
	}
	t.Errorf("Matrix %v not found in generated matrices", hasMatrix)
}

func TestReshapeToMatrix(t *testing.T) {
	// Test case: 1D slice to 2D matrix
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	res := reshapeToMatrix(input, 3)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Reshaped matrix is not as expected. Expected: %v, Got: %v", expected, res)
	}
}
